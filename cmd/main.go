package main

import (
	"log"
	"net/http"

	"github.com/atanda0x/bookstore/pkg/routes"
	"github.com/atanda0x/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
