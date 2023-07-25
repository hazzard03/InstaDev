package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("views/*html"))
	r := mux.NewRouter()

	r.PathPrefix("/assets/css").Handler(http.StripPrefix("/assets/css", http.FileServer(http.Dir("assets/css"))))

	r.PathPrefix("/assets/js").Handler(http.StripPrefix("/assets/js", http.FileServer(http.Dir("assets/js"))))

	r.HandleFunc("/login", indexHandler).Methods("GET")
	r.HandleFunc("/create-user", createUser).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "cadastro.html", nil)
}
