package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// Secure password prompt (https://www.mycodesmells.com/post/reading-password-input-in-go)
	fmt.Println("Enter password to check: ")
	userPassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	// Create sha1 hash in hex
	hash := sha1.New()
	hash.Write(userPassword)
	byteslice := hash.Sum(nil)
	hashedUserPassword := strings.ToUpper(hex.EncodeToString(byteslice))
	beginningOfHashedUserPassword := hashedUserPassword[:5]

	// Get all matching hashes from "have i been pwned".
	// https://haveibeenpwned.com/API/v2 (Searching by range)
	response, err := http.Get("https://api.pwnedpasswords.com/range/" + beginningOfHashedUserPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	passwordHashesInDatabase := strings.Split(string(body), "\r\n")

	// Check received hashes against full user password hash
	amount := checkForMatch(hashedUserPassword, passwordHashesInDatabase)
	if amount == "none" {
		fmt.Println("This password has not been found in any data breaches and is safe to use.")
	} else {
		fmt.Printf("This password was found %s times in data breaches.", amount)
		fmt.Println()
	}
}

func checkForMatch(hashedUserPassword string, passwordHashesInDatabase []string) string {
	for _, passwordHashInDatabase := range passwordHashesInDatabase {
		hashInDatabase := strings.Split(passwordHashInDatabase, ":")

		if hashedUserPassword[:5]+hashInDatabase[0] == hashedUserPassword {
			return hashInDatabase[1]
		}
	}
	return "none"
}
