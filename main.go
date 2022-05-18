package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getStaticHandler() http.Handler {
	dir := http.Dir("./static/")
	staticHandler := http.StripPrefix("/static/", http.FileServer(dir))
	return staticHandler
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods(http.MethodGet)
	r.HandleFunc("/produkt/{productId}", productHandler).Methods(http.MethodGet, http.MethodHead)
	r.PathPrefix("/static/").Handler(getStaticHandler())
	return r
}

func main() {
	r := getRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
