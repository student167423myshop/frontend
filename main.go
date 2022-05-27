package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type keySessionID struct{}

func getStaticHandler() http.Handler {
	dir := http.Dir("./static/")
	staticHandler := http.StripPrefix("/static/", http.FileServer(dir))
	return staticHandler
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods(http.MethodGet)
	r.HandleFunc("/produkt/{productId}", productHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/koszyk", viewCartHandler).Methods(http.MethodGet, http.MethodHead)
	//r.HandleFunc("/koszyk", addToCartHandler).Methods(http.MethodPost)
	//r.HandleFunc("/koszyk/empty", emptyCartHandler).Methods(http.MethodPost)
	r.PathPrefix("/static/").Handler(getStaticHandler())
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

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
