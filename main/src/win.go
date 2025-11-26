package main

import "fmt"

const (
	ligne  = 6
	Colone = 7
	Win    = 4
)

func checkVictory(grid [][]int, row, col int) bool {
	player := grid[row][col]
	if player == 0 {
		return false
	}

	count := 0
	for c := 0; c < Colone; c++ {
		if grid[ligne][c] == player {
			count++
			if count == Win {
				fmt.Printf("Victoire du joueur %d (horizontal)\n", player)
				return true
			}
		} else {
			count = 0
		}
	}

	count = 0
	for r := 0; r < ligne; r++ {
		if grid[r][col] == player {
			count++
			if count == Win {
				fmt.Printf("Victoire du joueur %d (vertical)\n", player)
				return true
			}
		} else {
			count = 0
		}
	}

	return false
}
func checkDraw(grid [][]int) bool {
	for r := 0; r < ligne; r++ {
		for c := 0; c < Colone; c++ {
			if grid[r][c] == 0 {
				return false
			}
		}
	}

	fmt.Println("Match nul ! La grille est pleine.")
	return true
}
