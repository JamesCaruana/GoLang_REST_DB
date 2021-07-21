package main

import (
	"context"
	"encoding/json"
	"go_rad/db"
	"go_rad/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = db.ConnectDb()

func main() {
	// Subrouter
	r := mux.NewRouter()

	// Create the handles for each
	r.HandleFunc("/api/users", getUsers).Methods(http.MethodGet)
	//r.HandleFunc("/api/users/{id}", getUserById).Methods(http.MethodGet)
	//r.HandleFunc("/api/user", createUser).Methods(http.MethodPut)
	//r.HandleFunc("/api/users/{id}", updateUser).Methods(http.MethodPut)
	//r.HandleFunc("/api/users/{id}", deleteUser).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:27017", r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Array of struct users
	var users []models.User

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		db.GetError(err, w)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users)
}

//func getUserById()

//func createUser()

//func updateUser()

//func deleteUser()
