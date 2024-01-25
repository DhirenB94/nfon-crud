package main

import (
	"log"
	"net/http"
	srv "nfon-crud/server"
	inMemDB "nfon-crud/storage"
)

func main() {
	inMemDb := inMemDB.NewInMemDB()
	server := srv.NewServer(inMemDb)
	err := http.ListenAndServe(":3333", server.Router)
	if err != nil {
		log.Fatal(err)
	}
}
