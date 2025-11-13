package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Parcourir tous les fichiers de templates dans le dossier templates (*.html) le * indique tous les fichiers
	listTemplates, errTemplates := template.ParseGlob("./templates/*.html")
	if errTemplates != nil {
		log.Fatal("Erreur lors du parsing des templates :", errTemplates)
	}

	// la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue au Puissance 4!"))
	})
	// Declarer une variable de type structure pour passer des données au templates
	type playData struct {
		Title string
		Tour  int
		Grid  [6][7]string
	}

	// route pour la page de jeu
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		data := playData{
			Title: "Puissance 4 - Partie en cours",
			Tour:  1,
			Grid:  [6][7]string{},
		}
		listTemplates.ExecuteTemplate(w, "play", data)
	})

	// declaration du serveur http
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
