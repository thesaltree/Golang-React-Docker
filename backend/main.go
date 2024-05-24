package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

func main() {
	initDBWithRetry()

	defer db.Close()

	fmt.Println("Successfully connected to the database!")

	// Start the router
	Routers()
}

func initDBWithRetry() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	for i := 0; i < 10; i++ { // Retry up to 10 times
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Error opening database: %v", err)
			time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
			continue
		}

		err = db.Ping()
		if err == nil {
			// Connection successful
			return
		}

		log.Printf("Error connecting to the database: %v", err)
		time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
	}

	// If still failing after retries, log the fatal error
	log.Fatalf("Could not connect to the database after multiple attempts: %v", err)
}

func Routers() {
	router := mux.NewRouter()
	router.HandleFunc("/users",
		GetUsers).Methods("GET")
	router.HandleFunc("/users",
		CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",
		GetUser).Methods("GET")
	router.HandleFunc("/users/{id}",
		UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",
		DeleteUser).Methods("DELETE")
	http.ListenAndServe(":3000",
		&CORSRouterDecorator{router})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	results, err := db.Query("Select id, first_name, middle_name, last_name, email, gender, civil_status, birthday, contact, address from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer results.Close()

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Gender, &user.CivilStatus, &user.Birthday, &user.Contact, &user.Address)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("insert into users (first_name, middle_name, last_name, email, gender, civil_status, birthday, contact, address) values (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var keyVal map[string]string
	err = json.Unmarshal(body, &keyVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	firstName := keyVal["firstName"]
	middleName := keyVal["middleName"]
	lastName := keyVal["lastName"]
	email := keyVal["email"]
	gender := keyVal["gender"]
	civilStatus := keyVal["civilStatus"]
	birthday := keyVal["birthday"]
	contact := keyVal["contact"]
	address := keyVal["address"]

	_, err = stmt.Exec(firstName, middleName, lastName, email, gender, civilStatus, birthday, contact, address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "New user was created")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	results, err := db.Query("select id, first_name, middle_name, last_name, email, gender, civil_status, birthday, contact, address from users where id= ?", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer results.Close()
	var user User
	userFound := false

	for results.Next() {
		err = results.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Gender, &user.CivilStatus, &user.Birthday, &user.Contact, &user.Address)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userFound = true
	}

	err = results.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found with ID: %s", params["id"])
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("update users set first_name=?, middle_name=?, last_name=?, email=?, gender=?, civil_status=?, birthday=?, contact=?, address=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var userUpdate User
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		panic(err.Error())
	}

	result, err := stmt.Exec(
		userUpdate.FirstName,
		userUpdate.MiddleName,
		userUpdate.LastName,
		userUpdate.Email,
		userUpdate.Gender,
		userUpdate.CivilStatus,
		userUpdate.Birthday,
		userUpdate.Contact,
		userUpdate.Address,
		params["id"],
	)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No user found with ID = %s", params["id"])
		return
	}

	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("delete from users where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No user found with ID = %s", params["id"])
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	CivilStatus string `json:"civilStatus"`
	Birthday    string `json:"birthday"`
	Contact     string `json:"contact"`
	Address     string `json:"address"`
}

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}

	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
