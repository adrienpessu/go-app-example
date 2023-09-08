package main

import (
	"context"
	"fmt"
	"github.com/adrienpessu/go-lib-example/github"
	"golang.org/x/crypto/ssh"
	"math/rand"
	"net"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	// Do nothing
	fmt.Println("Hello World!")

	// Call the function from the lib
	client := github.NewClient(nil)

	// list all organizations for user "adrienpessu"
	orgs, _, err := client.Organizations.List(context.Background(), "adrienpessu", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(orgs)
	fmt.Println(generatePassword())

	insecureIgnoreHostKey()
	insecureHostKeyCallback()
}

func generatePassword() string {
	s := make([]rune, 20)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}

func insecureIgnoreHostKey() {
	_ = &ssh.ClientConfig{
		User:            "username",
		Auth:            []ssh.AuthMethod{nil},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func insecureHostKeyCallback() {
	_ = &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{nil},
		HostKeyCallback: ssh.HostKeyCallback(
			func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			}),
	}
}
