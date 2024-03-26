package sudoku

import (
	"encoding/json"
	"log"
	"os"
)

const Dim int = 9

type Grid struct {
	values     [Dim][Dim]byte
	candidates [Dim][Dim][Dim]byte
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

	//grid.candidates[Dim-1][Dim-1] = [Dim]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}

	return &grid
}

func (g *Grid) CellIsEmpty(i, j int) bool {
	return g.values[i][j] == 0
}

func (g *Grid) Write(i, j int, b byte) {
	g.values[i][j] = b
	g.candidates[i][j] = [Dim]byte{}
}

func (g *Grid) Value(i, j int) byte {
	return g.values[i][j]
}

func (g *Grid) Candidates(i, j int) [Dim]byte {
	return g.candidates[i][j]
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
