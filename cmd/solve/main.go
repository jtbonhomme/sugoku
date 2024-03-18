package main

import (
	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

var g = sudoku.Grid{
	{9, 0, 0, 1, 0, 0, 0, 0, 5},
	{0, 0, 5, 0, 9, 0, 2, 0, 1},
	{8, 0, 0, 0, 4, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 0, 0},
	{0, 0, 0, 7, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 2, 6, 0, 0, 9},
	{2, 0, 0, 3, 0, 0, 0, 0, 6},
	{0, 0, 0, 2, 0, 0, 9, 0, 0},
	{0, 0, 0, 0, 0, 4, 5, 0, 0}}

func main() {
	g.Display()
	g.IsValid(0)
	g.Display()
}
