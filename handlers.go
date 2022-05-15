package main

import (
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

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
)
