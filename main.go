package main

import (
	"fmt"
	"log"
	"net/http"
	"restful/api"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/tasks", api.Get)
	handler.HandleFunc("/tasks/create", api.Create)
	handler.HandleFunc("/tasks/update", api.Update)
	handler.HandleFunc("/tasks/delete", api.Del)

	fmt.Println("Running on port 8080")
	log.Fatalln(http.ListenAndServe("localhost:8080", handler))
}
