package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// A Route is an endpoint route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds an array of the routes to pass to gorilla
type Routes []Route

// NewRouter create and return a new router
func (p *program) NewRouter() *mux.Router {
	// make routes
	var routes = Routes{
		Route{
			"SIL",
			"POST",
			"/sil/{silType}",
			p.makeSIL,
		},
		Route{
			"OBJ",
			"POST",
			"/obj",
			p.makeOBJ,
		},
		Route{
			"ADD",
			"POST",
			"/add/{silID}",
			p.add,
		},
		Route{
			"WRITE",
			"POST",
			"/write/{silID}",
			p.write,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
