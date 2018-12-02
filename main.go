package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//User : basic user structure
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

var users []User

//GetUserEndPoint : get a user by id
func GetUserEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, u := range users {
		if u.ID == params["id"] {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	//return empty User if nil
	json.NewEncoder(w).Encode(&User{})
}

//GetUsersEndPoint : gets a list of users
func GetUsersEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(users)
}

//CreateUserEndPoint : creates a user with an id
func CreateUserEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var u User
	_ = json.NewDecoder(req.Body).Decode(&u)
	u.ID = params["id"]
	users = append(users, u)
	json.NewEncoder(w).Encode(users)
}

//DeleteUserEndPoint : deletes a user by id
func DeleteUserEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, u := range users {
		if u.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	users = append(users, User{ID: "1", FirstName: "Ken", LastName: "L"})
	users = append(users, User{ID: "2", FirstName: "Jen", LastName: "Smith"})
	users = append(users, User{ID: "3", FirstName: "Jill", LastName: "Jones"})
	users = append(users, User{ID: "4", FirstName: "Joe", LastName: "James"})

	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsersEndPoint).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserEndPoint).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUserEndPoint).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUserEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9876", router))
}
