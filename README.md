# Golang-React-Docker

This repository contains a boilerplate project that integrates Golang as the backend, React as the frontend, and Docker for containerization. This setup is ideal for developers looking to build scalable and maintainable web applications with modern technologies.

## Features

- **Golang Backend**: A robust and efficient backend using Golang.
- **React Frontend**: A dynamic and responsive frontend using React.
- **Docker Integration**: Seamless containerization and orchestration using Docker.
- **Modular Architecture**: Clear separation of backend and frontend code for better maintainability.
- **Hot Reload**: Support for hot reloading in development mode.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker and Docker Compose installed on your machine.
- Node.js and npm (for frontend development).
- Go (Golang) installed on your machine.

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/thesaltree/Golang-React-Docker.git
cd Golang-React-Docker
```

### Running the Project

You can run the entire project using Docker Compose.

#### Build and start the containers:

```
docker-compose up --build
```

This command will build the Docker images and start the containers for both the Golang backend and React frontend.

#### Access the application:

Backend: http://localhost:8080 

Frontend: http://localhost:3000

### Development

#### Backend

To develop the backend locally without Docker, follow these steps:

- Navigate to the backend directory:
```
cd backend
```
- Install dependencies and run the server:
```
go mod tidy
go run main.go
```
The backend server will be available at http://localhost:8080.

#### Frontend

To develop the frontend locally without Docker, follow these steps:

- Navigate to the frontend directory:
```
cd frontend
```
- Install dependencies and start the development server:
```
npm install
npm start
```
The frontend server will be available at http://localhost:3000.

### Project Structure

```
.
├── backend             # Golang backend code
│   ├── main.go
│   ├── handlers
│   └── ...
├── frontend            # React frontend code
│   ├── public
│   ├── src
│   └── ...
├── docker-compose.yml  # Docker Compose configuration
└── README.md           # Project documentation
```

### Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes. Make sure to follow the coding standards and include relevant tests.

### License
This project is licensed under the MIT License - see the LICENSE file for details.


