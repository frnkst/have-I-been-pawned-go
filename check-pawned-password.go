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
	fmt.Println("Enter password to check: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	h := sha1.New()
	h.Write(bytePassword)
	bs := h.Sum(nil)
	hexHash := strings.ToUpper(hex.EncodeToString(bs))

	firstPart := hexHash[:5]

	response, err := http.Get("https://api.pwnedpasswords.com/range/" + firstPart)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	hashArray := strings.Split(string(body), "\n")

	for _, b := range hashArray {
		hashes := strings.Split(b, ":")

		if firstPart+hashes[0] == hexHash {
			fmt.Printf("This password was found %s times in data breaches", strings.TrimSuffix(hashes[1], "\r"))
			break
		}
	}
}
