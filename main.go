package main

import (
	"DevOps/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.Route()

	// Start the HTTP server on port 8080
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
