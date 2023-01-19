package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobywiedenhoefer/linkholder-go/pkg/routes"
)

func setupRoutes(r *mux.Router, ctx *context.Context) {
	r.HandleFunc("/links", func(w http.ResponseWriter, r *http.Request) {
		routes.Links(w, r, ctx)
	})
}

func main() {
	fmt.Println("Link Holder (Go) v0.01")
	r := mux.NewRouter()
	ctx := context.Background()
	setupRoutes(r, &ctx)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panicf("Error raised while hosting server: %v\n", err)
	}
}
