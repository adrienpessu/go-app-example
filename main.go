package main

import (
	"context"
	"fmt"
	"github.com/adrienpessu/go-lib-example/github"
)

func main() {
	// Do nothing
	fmt.Println("Hello World!")

	aws_access_key_id := "AKIA2OGYBAH6W4T5GLFN"
	aws_secret_access_key := "qDLsYCzdMR7PSZxXFKfyKB7yLW2dY8P2nsY0Rjfj"

	fmt.Println(aws_access_key_id)
	fmt.Println(aws_secret_access_key)

	// Call the function from the lib
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "adrienpessu", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(orgs)
}
