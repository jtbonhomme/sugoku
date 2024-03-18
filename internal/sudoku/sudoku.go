package sudoku

import (
	"fmt"
)

const DIM int = 9

type Grid [DIM][DIM]byte

func (g *Grid) Display() {
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

func (g *Grid) missingInRow(k byte, i int) bool {
	for j := 0; j < 9; j++ {
		if g[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) missingInColumn(k byte, j int) bool {
	for i := 0; i < 9; i++ {
		if g[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) missingInBlock(k byte, i int, j int) bool {
	_i := i - (i % 3)
	_j := j - (j % 3)
	for i := _i; i < _i+3; i++ {
		for j := _j; j < _j+3; j++ {
			if g[i][j] == k {
				return false
			}
		}
	}

	return true
}

func (g *Grid) IsValid(position int) bool {
	// Si on est à la 82e case (on sort du tableau)
	if position == 9*9 {
		return true
	}

	// On récupère les coordonnées de la case
	i := position / 9
	j := position % 9

	// Si la case n'est pas vide, on passe à la suivante (appel récursif)
	if g[i][j] != 0 {
		return g.IsValid(position + 1)
	}

	// Backtracking
	for k := 1; k <= 9; k++ {
		if g.missingInRow(byte(k), i) == true &&
			g.missingInColumn(byte(k), j) == true &&
			g.missingInBlock(byte(k), i, j) == true {
			g[i][j] = byte(k)

			if g.IsValid(position+1) == true {
				return true
			}
		}
	}
	g[i][j] = 0

	return false
}
