package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a function to generate passwords
func generatePassword(length int, useDigits bool, useSpecialChars bool) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if useDigits {
		charset += "0123456789"
	}
	if useSpecialChars {
		charset += "!@#$%^&*()-_=+/"
	}

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func main() {
	var length int // Declare length variable

	fmt.Print("Enter password length: ")
	_, err := fmt.Scan(&length) // Read password length input from user
	if err != nil || length <= 0 {
		fmt.Println("Invalid input. Using default password length of 12.")
		length = 12
	}

	useDigits := true       // Default to including digits in the password
	useSpecialChars := true // Default to including special characters in the password

	password := generatePassword(length, useDigits, useSpecialChars)
	fmt.Println("Generated Password:", password)
}
