package main

//:::: the route multiplexer lib is here
import "github.com/gorilla/mux"

// Create New Router with all routes  \\
func NewAPIRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true) // let's create the api router here

	// Let's register routes in our router
	for _, route := range api_routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	rsl := router

	return rsl // let's return the rooter
}
