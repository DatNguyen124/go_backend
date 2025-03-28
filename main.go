package main

import (
	"fmt"
	"log"
	"net/http"

	"go-bookstore/pkg/config"
	"go-bookstore/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	fmt.Println("✅ Connected to PostgreSQL database successfully!")

	// Setup router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// Start server
	fmt.Println("🚀 Server is running at http://localhost:8080")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
