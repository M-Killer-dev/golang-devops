package routes

import (
	"controllers"
	"net/http"
)

func Route() {
	// Register the handler function for the "/" route
	http.HandleFunc("/", controllers.Hello)
}
