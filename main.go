package main

import (
	"fmt"
	"log"
	"net/http"
	"yourapp/server"
)

func main() {
	fmt.Println("Project started")
	server.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
