package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
