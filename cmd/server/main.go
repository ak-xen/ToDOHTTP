package main

import (
	"log"
	"net/http"

	"github.com/ak-xen/ToDOHTTP/internal/task"
	transport "github.com/ak-xen/ToDOHTTP/transport/http"
)

func main() {

	repo := task.NewRepo()
	handler := transport.NewHandler(repo)
	router := transport.NewRouter(*handler)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
