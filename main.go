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
var error = template.Must(template.ParseFiles("views/pages/404.html"))

func hanldlerError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	error.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		hanldlerError(w, http.StatusInternalServerError)
	}
}

func main() {

	staticFiles := http.FileServer(http.Dir("assets"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "app", nil)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about", nil)
	})

	http.HandleFunc("/blog-post", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "blog-post", nil)
	})

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "blog", nil)
	})

	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "cart", nil)
	})

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "checkout", nil)
	})

	http.HandleFunc("/copmany", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "copmany", nil)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "contact", nil)
	})

	http.HandleFunc("/documentation", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "documentation", nil)
	})

	http.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "dummy", nil)
	})

	http.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "faq", nil)
	})

	http.HandleFunc("/fullpage", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "fullpage", nil)
	})

	http.HandleFunc("/pricing", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "pricing", nil)
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "product", nil)
	})

	http.HandleFunc("/real-state", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "real-state", nil)
	})

	http.HandleFunc("/restaurant", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "restaurant", nil)
	})

	http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "service", nil)
	})

	http.HandleFunc("/shop-item", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "shop-item", nil)
	})

	http.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "shop", nil)
	})

	http.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "sign-in", nil)
	})

	http.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "sign-up", nil)
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	log.Println("Ejecutando servidor en el puerto :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
