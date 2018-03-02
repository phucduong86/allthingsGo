// Bcrypt the input string
package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func BcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	//Todo: handle error
	return string(bytes), err
}

func CheckPasswordBcrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := os.Args[1]
	bcrypted, _ := BcryptPassword(password)
	//Todo: handle error
	fmt.Println("Password:", password)
	fmt.Println("Bcrypt:    ", bcrypted)

	match := CheckPasswordBcrypt(password, bcrypted)
	fmt.Println("Match:   ", match)
}
