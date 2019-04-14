# have-I-been-pawned-go
Safely check if your password has been pawned using go

## How it works

This program let's your enter a password, which then will be hashed. The first 5 letters of the hash will be sent to "have I been pawned"
(https://haveibeenpwned.com/API/v2). The api will return all password hashes that match those first 5 letters. The program will then go
through the list of passwords to see if the one you entered is found. If yes, it will tell you how many times the password occured in
databreaches.

## How to use it

### Build it yourself from the source
- Clone repository: `git clone https://github.com/frnkst/have-I-been-pawned-go.git`
- `go get golang.org/x/crypto/ssh/terminal`
- `go build check-pawned-password.go`
- `./check-pawned-password`

### Don't have go installed? Cross-compile it using docker 
- Clone repository: `git clone https://github.com/frnkst/have-I-been-pawned-go.git`
- Specify GOOS and GOARCH depending on your system (https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
- `docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=darwin -e GOARCH=amd64 golang:1.8 go get golang.org/x/crypto/ssh/terminal && go build -v`

### Download binary
- Download binary from releases
- Run the binary using ./check-pawned-password
