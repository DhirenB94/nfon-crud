package main

import (
	"log"
	"net/http"
	srv "nfon-crud/server"
)

func main() {
	server := srv.NewServer()
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		log.Fatal(err)
	}
}
