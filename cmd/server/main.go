package main

import (
	"log"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, Stepik!"))
}

func main() {
	http.HandlerFunc("/", router)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
