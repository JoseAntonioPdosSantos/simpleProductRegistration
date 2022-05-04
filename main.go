package main

import (
	"net/http"
	"simpleProductRegistration/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
