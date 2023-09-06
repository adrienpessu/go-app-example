package main

import (
	"context"
	"fmt"
	"github.com/adrienpessu/go-lib-example/github"
)

func main() {
	// Do nothing
	fmt.Println("Hello World!")

	// Call the function from the lib
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "adrienpessu", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(orgs)
}
