package main

import (
	"go_rad/db"
	"net/http"

	"github.com/gorilla/mux"
)

var collection = db.ConnectDb()

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/users", getUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{id}", getUserById).Methods(http.MethodGet)
	r.HandleFunc("/api/user", createUser).Methods(http.MethodPut)
	r.HandleFunc("/api/users/{id}", updateUser).Methods(http.MethodPut)
	r.HandleFunc("/api/users/{id}", deleteUser).Methods(http.MethodDelete)

}

func getUsers()

func getUserById()

func createUser()

func updateUser()

func deleteUser()
