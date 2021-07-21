package main

import (
	"context"
	"encoding/json"
	"go_rad/db"
	"go_rad/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = db.ConnectDb()

func main() {
	// Subrouter
	router := mux.NewRouter()

	// Create the handles for each
	router.HandleFunc("/api/users", getUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", getUserById).Methods(http.MethodGet)
	router.HandleFunc("/api/user", createUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}", updateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}", deleteUser).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:800", router))
}

// GET all users from database
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Array of struct users
	var users []models.User

	// Gets all user from database
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

// GET user by id
func getUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	// Filters by id
	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)

}

// POST new user to database
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	// Decodes json body to User struct
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Set the structure fields
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Inserts decoded data to database
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		db.GetError(err, w)
	}
	json.NewEncoder(w).Encode(res)

}

// PUT updates user by ID
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	var user models.User

	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Filters by id
	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Setting up structure
	update := bson.D{
		{"$set", bson.D{
			{"name", user.Name},
			{"dob", user.DOB},
			{"address", user.Address},
			{"description", user.Description},
			{"ca", user.CreatedAt},
			{"ua", time.Now()},
		}},
	}

	err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&user)
	if err != nil {
		db.GetError(err, w)
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}

// DELETE user by id from database
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Filters by id
	filter := bson.M{"_id": id}

	deleteRes, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteRes)
}
