package service

import (
	"fmt"
)

type AccountData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LogIn(data AccountData) {

	// Process the JSON data
	fmt.Println("Received JSON data:")
	fmt.Println("Login:", data.Login)
	fmt.Println("Password:", data.Password)
}
