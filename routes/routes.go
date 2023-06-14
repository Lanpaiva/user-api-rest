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
	r.HandleFunc("/users", controller.AllUsers)
	r.HandleFunc("/users/{id}", controller.FindById)

	defer log.Fatal(http.ListenAndServe(":8000", r))
}
