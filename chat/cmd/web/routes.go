package main

import (
	"github.com/RistoFlink/websockets-go/internal/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
