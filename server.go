package main

import (
	"log"
	"net/http"
)

func handler(req *http.Request) {
	username := req.URL.Query()["username"][0]
	log.Printf("user %s logged in.\n", username)
}
