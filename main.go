package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

func getStaticHandler() http.Handler {
	dir := http.Dir("./static/")
	staticHandler := http.StripPrefix("/static/", http.FileServer(dir))
	return staticHandler
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods(http.MethodGet)
	r.HandleFunc("/produkt/{productId}", productHandler).Methods(http.MethodGet)
	r.HandleFunc("/koszyk", viewCartHandler).Methods(http.MethodGet)
	r.HandleFunc("/koszyk", addToCartHandler).Methods(http.MethodPost)
	r.HandleFunc("/koszyk/usun", emptyCartHandler).Methods(http.MethodPost)
	r.HandleFunc("/zamowienie", newOrderHandler).Methods(http.MethodPost)
	r.HandleFunc("/zamowienie/{orderId}", viewOrderHandler).Methods(http.MethodGet)
	r.PathPrefix("/static/").Handler(getStaticHandler())
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

	return r
}
