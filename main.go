package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

const DIM int = 9
const DEBUG = false

type Grid [DIM][DIM]byte

var init_grille = Grid{
	{9, 0, 0, 1, 0, 0, 0, 0, 5},
	{0, 0, 5, 0, 9, 0, 2, 0, 1},
	{8, 0, 0, 0, 4, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 0, 0},
	{0, 0, 0, 7, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 2, 6, 0, 0, 9},
	{2, 0, 0, 3, 0, 0, 0, 0, 6},
	{0, 0, 0, 2, 0, 0, 9, 0, 0},
	{0, 0, 1, 9, 0, 4, 5, 7, 0}}

func dispGrid(g Grid) {
	for row := range g {
		fmt.Printf("\t")
		if row%3 == 0 {
			for i := 0; i < DIM+3; i++ {
				fmt.Printf("--")
			}
			fmt.Printf("-\n\t")
		}
		for col := range g[row] {
			if col%3 == 0 {
				fmt.Printf("| ")
			}
			if g[row][col] != 0 {
				fmt.Printf("%d ", g[row][col])
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("\t")
	for i := 0; i < DIM+3; i++ {
		fmt.Printf("--")
	}
	fmt.Printf("\n")
}

func absentSurLigne(k byte, grille Grid, i int) bool {
	if DEBUG {
		logger.Printf("is %d absent on row %d ?\n", k, i)
	}
	for j := 0; j < 9; j++ {
		if grille[i][j] == k {
			return false
		}
	}
	return true
}

func absentSurColonne(k byte, grille Grid, j int) bool {
	if DEBUG {
		logger.Printf("is %d absent on column %d ?\n", k, j)
	}
	for i := 0; i < 9; i++ {
		if grille[i][j] == k {
			return false
		}
	}
	return true
}

func absentSurBloc(k byte, grille Grid, i int, j int) bool {
	if DEBUG {
		logger.Printf("is %d absent in block (%d,%d) ?\n", k, i, j)
	}
	_i := i - (i % 3)
	_j := j - (j % 3)
	for i := _i; i < _i+3; i++ {
		for j := _j; j < _j+3; j++ {
			if grille[i][j] == k {
				return false
			}
		}
	}
	return true
}

func estValide(grille *Grid, position int) bool {
	if DEBUG {
		dispGrid(*grille)
	}
	// Si on est à la 82e case (on sort du tableau)
	if position == 9*9 {
		return true
	}

	// On récupère les coordonnées de la case
	i := position / 9
	j := position % 9

	// Si la case n'est pas vide, on passe à la suivante (appel récursif)
	if grille[i][j] != 0 {
		return estValide(grille, position+1)
	}
	if DEBUG {
		fmt.Printf("Position: %d (%d,%d)\n", position, i, j)
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	// Backtracking
	for k := 1; k <= 9; k++ {
		if absentSurLigne(byte(k), *grille, i) == true && absentSurColonne(byte(k), *grille, j) == true && absentSurBloc(byte(k), *grille, i, j) == true {
			grille[i][j] = byte(k)

			if estValide(grille, position+1) == true {
				return true
			}
		}
	}
	grille[i][j] = 0

	return false
}

func main() {
	dispGrid(init_grille)
	estValide(&init_grille, 0)
	dispGrid(init_grille)
}
