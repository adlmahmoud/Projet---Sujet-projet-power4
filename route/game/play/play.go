package main

import (
	"fmt"
	"net/http"
)

func playHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="fr">
<head>
  <meta charset="UTF-8">
  <title>Puissance 4</title>
  <link rel="stylesheet" href="style.css">
</head>
<body>

<h1>Puissance 4</h1>

<div id="current-player">Joueur 1 (Rouge), à vous de jouer !</div>

<table border="1" cellspacing="0" cellpadding="10">`)
	for lin := 0; lin < 6; lin++ {
		fmt.Fprintf(w, "<tr>")
		for col := 0; col < 7; col++ {
			fmt.Fprintf(w, `<td data-row="%d" data-col="%d"></td>`, lin, col)
		}
		fmt.Fprintf(w, "</tr>")
	}
	fmt.Fprintf(w, `</table>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", playHandler)
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
