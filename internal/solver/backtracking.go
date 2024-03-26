package solver

import (
	"time"

	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func SolveWithBacktracking(position int, grid *sudoku.Grid) bool {
	time.Sleep(time.Millisecond * 7)
	// Si on est à la 82e case (on sort du tableau)
	if position == 9*9 {
		return true
	}

	// On récupère les coordonnées de la case
	i := position / 9
	j := position % 9

	// Si la case n'est pas vide, on passe à la suivante (appel récursif)
	if grid[i][j] != 0 {
		return SolveWithBacktracking(position+1, grid)
	}

	// Backtracking
	for k := 1; k <= 9; k++ {
		if grid.MissingInRow(byte(k), i) &&
			grid.MissingInColumn(byte(k), j) &&
			grid.MissingInBlock(byte(k), i, j) {
			grid[i][j] = byte(k)

			if SolveWithBacktracking(position+1, grid) {
				return true
			}
		}
	}
	grid[i][j] = 0

	return false
}
