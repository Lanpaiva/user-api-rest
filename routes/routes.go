package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanpaiva/teste/controller"
)

func Handler() {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/users", controller.AllUsers).Methods("Get")
	r.HandleFunc("/users/{id}", controller.FindById).Methods("Get")
	r.HandleFunc("/users", controller.NewUser).Methods("Post")
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods("Delete")
	r.HandleFunc("/users/{id}", controller.UpdateUser).Methods("Put")

	defer log.Fatal(http.ListenAndServe(":8000", r))
}
