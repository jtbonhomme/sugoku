package sudoku

import (
	"encoding/json"
	"log"
	"os"
)

const DIM int = 9

type Grid struct {
	Values [DIM][DIM]byte
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
	err = json.Unmarshal(fileBytes, &grid.Values)
	if err != nil {
		log.Fatal(err)
	}

	return &grid
}

func (g *Grid) MissingInRow(k byte, i int) bool {
	for j := 0; j < 9; j++ {
		if g.Values[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) MissingInColumn(k byte, j int) bool {
	for i := 0; i < 9; i++ {
		if g.Values[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) MissingInBlock(k byte, i int, j int) bool {
	_i := i - (i % 3)
	_j := j - (j % 3)
	for i := _i; i < _i+3; i++ {
		for j := _j; j < _j+3; j++ {
			if g.Values[i][j] == k {
				return false
			}
		}
	}

	return true
}
