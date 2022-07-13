package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "subrat"
	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:", hash)

	match := CheckPasswordAndHash(password, hash)
	fmt.Println("Match", match)
}
