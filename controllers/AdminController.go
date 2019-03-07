package controllers

import (
	"html/template"
	"net/http"
)

// AdminHome index page admin
func AdminHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/Admin/Index.html", "views/Admin/Partials/Header.html", "views/Admin/Partials/Footer.html")
	res := map[string]interface{}{
		"Title": "Admin Home",
	}
	t.Execute(w, res)
}
