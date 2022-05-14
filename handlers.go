package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"products": getProducts(),
	}); err != nil {
		panic(err.Error())
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := getProducts()
	product := products
	data, _ := json.Marshal(product)
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
)
