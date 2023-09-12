package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/akhil/go-bookstore/pkg/routes" // Importe routes a partir do caminho raiz
)



func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
