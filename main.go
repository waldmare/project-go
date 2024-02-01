package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Ultimate Password Generator!")

	useSpecialChars := askYesNo("Include special characters?")
	useUppercase := askYesNo("Include uppercase letters?")
	useDigits := askYesNo("Include digits?")
	useLowercase := !useUppercase || askYesNo("Include lowercase letters?")
	useMemorableWord := askYesNo("Include a memorable word in the password?")
	passwordLength := askForLength()

	password := generatePassword(useLowercase, useUppercase, useDigits, useSpecialChars, useMemorableWord, passwordLength)

	fmt.Println("\nYour ultimate generated password is:", password)
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

func generatePassword(useLowercase, useUppercase, useDigits, useSpecialChars, useMemorableWord bool, length int) string {
	characters := ""
	if useLowercase {
		characters += "abcdefghijklmnopqrstuvwxyz"
	}
	if useUppercase {
		characters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useDigits {
		characters += "0123456789"
	}
	if useSpecialChars {
		characters += "!@#$%^&*()-_=+[]{},.<>/?;:'\"\\|`~"
	}

	if characters == "" {
		fmt.Println("Please include at least one character set. Exiting.")
		return ""
	}

	password := make([]byte, length)

	if useMemorableWord {
		memorableWord := generateMemorableWord()
		copy(password, memorableWord)
	}

	for i := len(password) - 1; i >= 0; i-- {
		if password[i] == 0 {
			password[i] = characters[rand.Intn(len(characters))]
		}
	}

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func generateMemorableWord() []byte {
	memorableWords := []string{
		"sunshine", "giraffe", "butterfly", "rainbow", "waterfall",
		"whisper", "chocolate", "serendipity", "harmony", "moonlight",
	}

	word := []byte(memorableWords[rand.Intn(len(memorableWords))])

	if unicode.IsLower(rune(word[0])) {
		word[0] = byte(unicode.ToUpper(rune(word[0])))
	}

	return word
}
