package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}<>?/|"
)

func main() {
	var length int
	var lower, upper, numbers, symbols string

	fmt.Println("Length of the password: ")
	fmt.Scan(&length)

	fmt.Println("Do you want to have lowercase letters? (y/n)")
	fmt.Scan(&lower)

	fmt.Println("Do you want to have uppercase letters? (y/n)")
	fmt.Scan(&upper)

	fmt.Println("Do you want to have numbers? (y/n)")
	fmt.Scan(&numbers)

	fmt.Println("Do you want to have symbols (y/n)")
	fmt.Scan(&symbols)

	password, ok := generatePassword(
		length,
		strings.ToLower(lower) == "y",
		strings.ToLower(upper) == "y",
		strings.ToLower(numbers) == "y",
		strings.ToLower(symbols) == "y",
	)

	if ok {
		fmt.Println("Generated password: ", password)
	}
}

func generatePassword(length int, useLower, useUpper, useNumbers, useSymbols bool) (string, bool) {
	var charset string

	if useLower {
		charset += lowercase
	}
	if useUpper {
		charset += uppercase
	}
	if useNumbers {
		charset += numbers
	}
	if useSymbols {
		charset += symbols
	}

	if len(charset) == 0 || length < 1 {
		fmt.Println("There is no charset or length to make a password")
		return "", false
	}

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		char, err := getRandomChar(charset)
		if err != nil {
			fmt.Println("Error generating random character: ", err)
			return "", false
		}
		password[i] = char
	}

	return string(password), true
}

func getRandomChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}
