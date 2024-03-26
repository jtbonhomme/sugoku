package solver

import (
	"time"

	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func SolveWithBacktracking(position int, grid *sudoku.Grid, speed int) bool {
	time.Sleep(time.Millisecond * time.Duration(speed))
	// Si on est à la 82e case (on sort du tableau)
	if position == sudoku.Dim*sudoku.Dim {
		return true
	}

	// On récupère les coordonnées de la case
	row := position / sudoku.Dim
	col := position % sudoku.Dim

	// Si la case n'est pas vide, on passe à la suivante (appel récursif)
	if !grid.CellIsEmpty(row, col) {
		return SolveWithBacktracking(position+1, grid, speed)
	}

	// Backtracking
	for k := 1; k <= sudoku.Dim; k++ {
		if grid.MissingInRow(byte(k), row) &&
			grid.MissingInColumn(byte(k), col) &&
			grid.MissingInBlock(byte(k), row, col) {
			grid.Write(row, col, byte(k))

			if SolveWithBacktracking(position+1, grid, speed) {
				return true
			}
		}
	}
	grid.Write(row, col, 0)

	return false
}

func FillCandidatesAtPos(position int, grid *sudoku.Grid) {
	// On récupère les coordonnées de la case
	row := position / sudoku.Dim
	col := position % sudoku.Dim

	// Si la case n'est pas vide, on passe à la suivante (appel récursif)
	if !grid.CellIsEmpty(row, col) {
		return
	}
	//grid.ResetCandidates(row, col)

	for k := 1; k <= sudoku.Dim; k++ {
		if grid.MissingInRow(byte(k), row) &&
			grid.MissingInColumn(byte(k), col) &&
			grid.MissingInBlock(byte(k), row, col) {
			grid.SetCandidate(row, col, byte(k))
		}
	}
}

func FillCandidates(grid *sudoku.Grid) {
	for position := 0; position < sudoku.Dim*sudoku.Dim; position++ {
		// On récupère les coordonnées de la case
		row := position / sudoku.Dim
		col := position % sudoku.Dim

		// Si la case n'est pas vide, on passe à la suivante (appel récursif)
		if !grid.CellIsEmpty(row, col) {
			continue
		}
		//grid.ResetCandidates(row, col)
		for k := 1; k <= sudoku.Dim; k++ {
			if grid.MissingInRow(byte(k), row) &&
				grid.MissingInColumn(byte(k), col) &&
				grid.MissingInBlock(byte(k), row, col) {
				grid.SetCandidate(row, col, byte(k))
			}
		}
	}
}
