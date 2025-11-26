package main

type Partie struct {
	Grille        [6][7]rune
	Joueur1       rune
	Joueur2       rune
	JoueurCourant rune
	NombreDeTours int
	EtatDuJeu     string
}

func InitialiserPartie() Partie {
	var grille [6][7]rune

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			grille[i][j] = '.'
		}
	}

	joueur1 := 'X'
	joueur2 := 'O'

	partie := Partie{
		Grille:        grille,
		Joueur1:       joueur1,
		Joueur2:       joueur2,
		JoueurCourant: joueur1,
		NombreDeTours: 0,
		EtatDuJeu:     "en cours",
	}

	return partie
}
