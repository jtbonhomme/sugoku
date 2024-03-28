package sudoku

import (
	"encoding/json"
	"log"
	"os"
)

const Dim int = 9

type Grid struct {
	values     [Dim][Dim]byte      // rows x cols
	candidates [Dim][Dim][Dim]byte // rows x cols x candidates
}

func New() *Grid {
	return &Grid{}
}

func NewFromFile(filename string) *Grid {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := Grid{}
	err = json.Unmarshal(fileBytes, &grid.values)
	if err != nil {
		log.Fatal(err)
	}

	return &grid
}

func (g *Grid) CellIsEmpty(i, j int) bool {
	return g.values[i][j] == 0
}

func (g *Grid) Write(i, j int, b byte) {
	g.values[i][j] = b
	g.ResetCandidates(i, j)
}

func (g *Grid) Value(i, j int) byte {
	return g.values[i][j]
}

func (g *Grid) Candidates(i, j int) [Dim]byte {
	return g.candidates[i][j]
}

func (g *Grid) SetCandidate(i, j int, b byte) {
	if b < 1 || b > 9 {
		return
	}
	g.candidates[i][j][b-1] = b
}

func (g *Grid) ResetCandidates(i, j int) {
	g.candidates[i][j] = [Dim]byte{}
}

func (g *Grid) UnsetCandidate(i, j int, b byte) {
	if b < 1 || b > 9 {
		return
	}
	g.candidates[i][j][b-1] = 0
}

func (g *Grid) MissingInRow(k byte, i int) bool {
	for j := 0; j < 9; j++ {
		if g.values[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) MissingInColumn(k byte, j int) bool {
	for i := 0; i < 9; i++ {
		if g.values[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) MissingInBlock(k byte, i, j int) bool {
	_i := i - (i % 3)
	_j := j - (j % 3)
	for i := _i; i < _i+3; i++ {
		for j := _j; j < _j+3; j++ {
			if g.values[i][j] == k {
				return false
			}
		}
	}

	return true
}

func (g *Grid) IsRowComplete(row int) bool {
	result := [Dim]byte{}

	for col := 0; col < Dim; col++ {
		v := g.Value(row, col)
		if v == 0 {
			continue
		}

		result[v-1] = v
	}

	for i := 0; i < Dim; i++ {
		if result[i] != byte(i+1) {
			return false
		}
	}

	return true
}

func (g *Grid) IsColComplete(col int) bool {
	result := [Dim]byte{}

	for row := 0; row < Dim; row++ {
		v := g.Value(row, col)
		if v == 0 {
			continue
		}

		result[v-1] = v
	}

	for i := 0; i < Dim; i++ {
		if result[i] != byte(i+1) {
			return false
		}
	}

	return true
}

func (g *Grid) IsBlockComplete(row, col int) bool {
	row = row - row%3
	col = col - col%3
	result := [Dim]byte{}

	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			v := g.Value(r, c)
			if v == 0 {
				continue
			}

			result[v-1] = v
		}
	}

	for i := 0; i < Dim; i++ {
		if result[i] != byte(i+1) {
			return false
		}
	}

	return true
}

func (g *Grid) IsComplete() bool {
	for row := 0; row < Dim; row++ {
		if !g.IsRowComplete(row) {
			return false
		}

		for col := 0; col < Dim; col++ {
			if !g.IsColComplete(col) || !g.IsBlockComplete(row, col) {
				return false
			}
		}
	}

	return true
}
