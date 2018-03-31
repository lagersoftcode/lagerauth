package main

import (
	"bytes"
	"fmt"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Println("Enter password")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))

	fmt.Println("Enter password (again)")
	passConfirm, _ := terminal.ReadPassword(int(syscall.Stdin))

	if !bytes.Equal(password, passConfirm) {
		fmt.Println("Error: passwords doesnt match!")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashed Password: %s\n", string(hashedPassword))
}
