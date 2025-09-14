package main

import (
	"log"
	"net/http"
)

func main() {

	repo := task.NewRepository()
	handler := transport.NewHandler(repo)
	router := transport.NewRouter(handler)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
