package controllers

import (
	"html/template"
	"net/http"
)

type Vnexpress struct {
	Title string
	Href  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/Index.html")
	t.Execute(w, "")
}
