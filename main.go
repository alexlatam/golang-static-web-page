package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
}

var templates = template.Must(template.New("T").ParseGlob("views/**/*.html"))
var error = template.Must(template.ParseFiles("views/pages/error.html"))

func hanldlerError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	error.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// hanldlerError(w, http.StatusInternalServerError)
	}
}

func main() {

	staticFiles := http.FileServer(http.Dir("assets"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{Name: "Juan"}
		renderTemplate(w, "user", user)
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	log.Println("Ejecutando servidor en el puerto :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
