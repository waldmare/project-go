package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Password Generator")

	useSpecialChars := askYesNo("Include special characters?")
	useUppercase := askYesNo("Include uppercase letters?")
	passwordLength := askForLength()

	password := generatePassword(useSpecialChars, useUppercase, passwordLength)

	fmt.Println("\nYour generated password is:", password)
}

func askYesNo(question string) bool {
	var response string
	fmt.Printf("%s (y/n): ", question)
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}

func askForLength() int {
	var length int
	fmt.Print("Enter password length: ")
	fmt.Scanln(&length)
	return length
}

func generatePassword(useSpecialChars, useUppercase bool, length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	if useUppercase {
		characters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useSpecialChars {
		characters += "!@#$%^&*()-_=+[]{},.<>/?;:'\"\\|`~"
	}
	characters += "0123456789"

	password := make([]byte, length)
	for i := range password {
		password[i] = characters[rand.Intn(len(characters))]
	}

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}
