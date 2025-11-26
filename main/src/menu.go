package main

type Player1 struct {
	Name   string
	Jeton  string
	Grille [6][7]string
}

func (player *Player1) InitPlayer1() {
	*player = Player1{
		Name:  "Player1",
		Jeton: "X",
	}
}

type Player2 struct {
	Name   string
	Jeton  string
	Grille [6][7]string
}

func (player *Player2) InitPlayer2() {
	*player = Player2{
		Name:  "Player2",
		Jeton: "O",
	}
}
