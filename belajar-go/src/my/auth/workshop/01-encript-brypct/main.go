package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashedStr, err := encryptToBcrypt("hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashed String = %s", hashedStr)
}

// encrypt string kedalam bcrypt
func encryptToBcrypt(str string) (string, error) {
	// Task: Hashing the password with the default cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(str), 10)

	if err != nil {
		return "Invalid Encrypt to Bcrypt", err
	}

	return string(hash), nil // TODO: replace this
}
