package main

import (
	"net/http"

	"diception/routes"
)

func main() {
	// Set the route
	http.HandleFunc("/omg", routes.OMG)

	// Listen
	http.ListenAndServe(":8000", nil)
}
