package main

import (
	"context"
	"fmt"
	"github.com/adrienpessu/go-lib-example/github"
	"os"
)

func main() {
	// Do nothing
	fmt.Println("Hello World!")

	azureStorageKey := os.Getenv("AZURE_STORAGE_KEY")

	fmt.Println(azureStorageKey)

	// Call the function from the lib
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "adrienpessu", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(orgs)
}
