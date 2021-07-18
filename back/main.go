package main

import (
	"back/handlers"
	"back/repositories"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handlers.Init()
	repositories.Init()

	r := mux.NewRouter()
	r.HandleFunc("/pages", handlers.CreateNewPageHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/pages", handlers.ListAllPages).Methods(http.MethodGet, http.MethodOptions)

	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedHeaders([]string{"content-type"}),
		gorillaHandlers.AllowedOrigins([]string{"*"}),
		gorillaHandlers.AllowCredentials(),
	)

	r.Use(cors)

	log.Fatal(http.ListenAndServe(":3000", r))
}
