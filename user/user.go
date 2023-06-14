package user

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Senha string `json:"Senha"`
}

var Users []User
