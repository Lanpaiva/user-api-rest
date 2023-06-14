package main

import (
	"fmt"

	"github.com/lanpaiva/teste/database"
	"github.com/lanpaiva/teste/routes"
	"github.com/lanpaiva/teste/user"
)

func main() {
	user.Users = []user.User{
		{ID: 1, Nome: "user1", Senha: "123"},
		{ID: 2, Nome: "user2", Senha: "123"},
	}
	database.ConnectDB()

	fmt.Println("Server inicializado")
	routes.Handler()
}
