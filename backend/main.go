package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Routers()
}

func Routers() {
	InitDB()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/users",
		GetUsers).Methods("GET")
	router.HandleFunc("/users",
		CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",
		GetUser).Methods("GET")
	router.HandleFunc("/users/{id}",
		UpdateUser).Methods("PUT")
	//router.HandleFunc("/users/{id}",
	//	DeleteUser).Methods("DELETE")
	http.ListenAndServe(":3000",
		&CORSRouterDecorator{router})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	results, err := db.Query("Select id, first_name, middle_name, last_name, email, gender, civil_status, birthday, contact, address, age from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer results.Close()

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Gender, &user.CivilStatus, &user.Birthday, &user.Contact, &user.Address, &user.Age)
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
	stmt, err := db.Prepare("insert into users (id, first_name, middle_name, last_name, email, gender, civil_status, birthday, contact, address, age) values (?,?,?,?,?,?,?,?,?,?,?)")
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
	middleName := keyVal["middle_name"]
	lastName := keyVal["last_name"]
	email := keyVal["email"]
	gender := keyVal["gender"]
	civilStatus := keyVal["civil_status"]
	birthday := keyVal["birthday"]
	contact := keyVal["contact"]
	address := keyVal["address"]
	age := keyVal["age"]

	_, err = stmt.Exec(firstName, middleName, lastName, email, gender, civilStatus, birthday, contact, address, age)
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
	results, err := db.Query("select id, first_name, middle_name, last_name, email, gender, civil_status, birthday, contact, address, age from users where id= ?", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer results.Close()
	var user User
	userFound := false

	for results.Next() {
		err = results.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Gender, &user.CivilStatus, &user.Birthday, &user.Contact, &user.Address, &user.Age)
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
	stmt, err := db.Prepare("update users set first_name=?, middle_name=?, last_name=?, email=?, gender=?, civil_status=?, birthday=?, contact=?, address=?, age=? where id=?")
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

// Task 7: Write code for delete user here

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	CivilStatus string `json:"civilStatus"`
	Birthday    string `json:"birthday"`
	Contact     string `json:"contact"`
	Address     string `json:"address"`
	Age         string `json:"age"`
}

var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		panic(err.Error())
	}
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
