// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	password := "admin123" // Change this to your desired password
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Hashed password:", string(hash))
// }