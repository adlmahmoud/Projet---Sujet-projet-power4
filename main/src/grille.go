package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Lignes   = 6
	Colonnes = 7
)

type Grille [Lignes][Colonnes]rune

func NouvelleGrille() Grille {
	var g Grille
	for i := 0; i < Lignes; i++ {
		for j := 0; j < Colonnes; j++ {
			g[i][j] = '.'
		}
	}
	return g
}

func AfficherGrille(g Grille) {
	fmt.Print(" ")
	for col := 0; col < Colonnes; col++ {
		fmt.Printf(" %d", col)
	}
	fmt.Println()
	for i := 0; i < Lignes; i++ {
		for j := 0; j < Colonnes; j++ {
			fmt.Printf(" %c", g[i][j])
		}
		fmt.Println()
	}
}

func JouerCoup(g *Grille, col int, joueur rune) (int, bool) {
	if col < 0 || col >= Colonnes {
		return -1, false
	}
	for ligne := Lignes - 1; ligne >= 0; ligne-- {
		if g[ligne][col] == '.' {
			g[ligne][col] = joueur
			return ligne, true
		}
	}
	return -1, false
}

func VerifierVictoire(g Grille, ligne, col int, joueur rune) bool {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}

	for _, dir := range directions {
		count := 1

		for i := 1; i < 4; i++ {
			r := ligne + dir[0]*i
			c := col + dir[1]*i
			if r >= 0 && r < Lignes && c >= 0 && c < Colonnes && g[r][c] == joueur {
				count++
			} else {
				break
			}
		}
		for i := 1; i < 4; i++ {
			r := ligne - dir[0]*i
			c := col - dir[1]*i
			if r >= 0 && r < Lignes && c >= 0 && c < Colonnes && g[r][c] == joueur {
				count++
			} else {
				break
			}
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

func ChangerJoueur(joueur rune) rune {
	if joueur == 'X' {
		return 'O'
	}
	return 'X'
}

func main() {
	grille := NouvelleGrille()
	joueur := 'X'
	reader := bufio.NewReader(os.Stdin)

	for {
		AfficherGrille(grille)
		fmt.Printf("Joueur %c, entrez une colonne (0-%d) : ", joueur, Colonnes-1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		col, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Entrée invalide. Veuillez entrer un numéro de colonne.")
			continue
		}

		ligne, ok := JouerCoup(&grille, col, joueur)
		if !ok {
			fmt.Println("Colonne invalide ou pleine. Essayez encore.")
			continue
		}

		if VerifierVictoire(grille, ligne, col, joueur) {
			AfficherGrille(grille)
			fmt.Printf("🎉 Joueur %c a gagné !\n", joueur)
			break
		}

		grillePleine := true
		for c := 0; c < Colonnes; c++ {
			if grille[0][c] == '.' {
				grillePleine = false
				break
			}
		}
		if grillePleine {
			AfficherGrille(grille)
			fmt.Println("Match nul ! La grille est pleine.")
			break
		}

		joueur = ChangerJoueur(joueur)
	}
}
