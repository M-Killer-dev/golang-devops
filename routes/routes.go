package routes

import (
	"DevOps/controllers"
	"net/http"
)

func Route() {
	// Register the handler function for the "/" route
	http.HandleFunc("/", controllers.RootHandler)
	http.HandleFunc("/v1/tools/validate", controllers.ValidateHandler)
	http.HandleFunc("/v1/tools/lookup", controllers.LookupHandler)
	http.HandleFunc("/v1/history", controllers.HistoryHandler)
}
