package main

import (
	"log"
	"time"

	"github.com/jtbonhomme/sugoku/internal/solver"
	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func main() {
	var speed int

	grid := sudoku.New()
	solver.FillCandidates(grid)

	start := time.Now()
	log.Println("start backtracking sudoku")
	solver.SolveWithBacktracking(0, grid, speed)
	log.Printf("sudoku solved in %s", time.Since(start))
	log.Println("press CTRL+C to quit")

	log.Printf("\n%s\n", grid.ValuesToString())
	log.Printf("is grid correct: %v", grid.IsComplete())
	log.Println("Exit")
}
