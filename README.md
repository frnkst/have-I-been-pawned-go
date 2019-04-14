# have-I-been-pawned-go
Safely check if your password has been pawned using go

## How it works

This program let's your enter a password, which then will be hashed. The first 5 letters of the hash will be sent to have I been pawned
(https://haveibeenpwned.com/API/v2). The api will return all password hashes that match those first 5 letters. The program will then go
through the list of passwords to see if the one you entered is found. If yes, it will tell you how many times the password occured in
databreaches.

## How to use it

### Build from repo
- Clone repo
- go build check-pawned-password.go
- ./check-pawned-password

### Download binary
- Download binary from releases
- Run the binary using ./check-pawned-password
