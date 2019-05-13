package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ScanDestination",
		"POST",
		"/scan",
		ScanDestination,
	},
}

// NewRouter creates and returns a new request router.
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Logger logs when we receive requests, and the execution time of handling
// these requests.  We don't log client IP addresses or the given obfs4
// parameters.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// main is the entry point of this tool.
func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s CERT_FILE KEY_FILE\n", os.Args[0])
		os.Exit(1)
	}
	certFile := os.Args[1]
	keyFile := os.Args[2]

	router := NewRouter()
	log.Fatal(http.ListenAndServeTLS(":8080", certFile, keyFile, router))
}
