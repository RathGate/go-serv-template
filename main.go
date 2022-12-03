package main

import (
	"fmt"
	"go-serv-template/packages/tuto"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tasklist = tuto.Tasklist{
	Name: "Parties du tutoriel:",
	Tasks: []tuto.Task{
		{Name: "Créer le serveur", Done: true},
		{Name: "Lier un template html", Done: true},
		{Name: "Gérer les fichiers statiques", Done: true},
		{Name: "Envoyer des données au template", Done: false},
		{Name: "Utiliser plusieurs templates", Done: false},
	},
}

var person tuto.Person

func main() {

	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Gère la route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		pages := []string{
			"templates/index.html",
			"templates/greetings.html",
			"templates/tasklist.html",
		}
		tmpl := template.Must(template.ParseFiles(pages...))
		tmpl.Execute(w, tasklist)
	})

	// Gère la route "/form"
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {

		// r = request ! On peut récupérer toutes les valeurs de la requête:
		if r.Method == http.MethodPost {
			// Si un champ n'existe pas dans le form, il renverra "" en guise de valeur.
			// Ici, si le champ qui porte l'attribut name="name" n'est pas vide, le form a été rempli:
			if r.FormValue("name") != "" {
				person.Name = r.FormValue("name")
				person.Age, _ = strconv.Atoi(r.FormValue("age"))
				if r.FormValue("go") == "likes" {
					person.LikesGo = true
				} else {
					person.LikesGo = false
				}
				if r.FormValue("network") == "likes" {
					person.LikesNetwork = true
				} else {
					person.LikesNetwork = false
				}
			} else if r.FormValue("reset") == "reset" {
				person = tuto.Person{}
			}
		}

		// Si la page charge pour la première fois et que la requête n'est pas de type "post",
		// les lignes au dessus ne chargeront pas, mais le template oui ! ↓
		tmpl := template.Must(template.ParseFiles("templates/form.html"))
		tmpl.Execute(w, person)
	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
