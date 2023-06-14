package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lanpaiva/teste/database"

	"github.com/gorilla/mux"
	"github.com/lanpaiva/teste/user"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pág. Principal"))
	w.WriteHeader(http.StatusOK)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var client []user.User
	database.DB.Find(&client)
	json.NewEncoder(w).Encode(client)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, user := range user.Users {
		if strconv.Itoa(user.ID) == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}
