package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*.html"))
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":2025", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		executeTemplate(w, "login.html", users)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")

		found := false
		var index int

		for i, user := range users {
			if user.Email == email {
				found = true
				index = i
			}
		}

		if !found {
			executeTemplate(w, "notFound.html", nil)
			return
		}

		executeTemplate(w, "home.html", users[index])
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func executeTemplate(w http.ResponseWriter, name string, data any) {
	err := tpl.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
