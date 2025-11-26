package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AjouterPion(g *Grille, joueur rune) (int, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Joueur %c, entrez une colonne (0-%d) : ", joueur, Colonnes-1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		col, err := strconv.Atoi(input)

		if err != nil || col < 0 || col >= Colonnes {
			fmt.Println("❌ Colonne invalide. Veuillez entrer un numéro valide.")
			continue
		}

		for ligne := Lignes - 1; ligne >= 0; ligne-- {
			if g[ligne][col] == '.' {
				g[ligne][col] = joueur
				return ligne, col
			}
		}

		fmt.Println("⚠️ Colonne pleine. Choisissez une autre colonne.")
	}
}
