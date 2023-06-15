package controller

import (
	"encoding/json"
	"net/http"

	"github.com/lanpaiva/teste/database"

	"github.com/gorilla/mux"
	"github.com/lanpaiva/teste/user"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pág. Principal"))
	w.WriteHeader(http.StatusOK)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var delUser []user.User

	database.DB.Delete(&delUser, id)
	json.NewEncoder(w).Encode(delUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var upUser user.User

	database.DB.First(&upUser, id)
	json.NewDecoder(r.Body).Decode(&upUser)
	database.DB.Save(&upUser)
	json.NewEncoder(w).Encode(upUser)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var client []user.User
	database.DB.Find(&client)
	json.NewEncoder(w).Encode(client)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var client []user.User

	database.DB.First(&client, id)
	json.NewEncoder(w).Encode(client)
}
