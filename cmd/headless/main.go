package main

import (
	"flag"
	"log"
	"time"

	"github.com/jtbonhomme/sugoku/internal/solver"
	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func main() {
	var speed int
	var filename string

	flag.StringVar(&filename, "f", "_examples/empty.json", "initial grid (default is _examples/backtracking.json)")
	flag.Parse()

	grid := sudoku.NewFromFile(filename)

	start := time.Now()
	log.Println("start backtracking sudoku")
	solver.SolveWithBacktracking(0, grid, speed, true)
	log.Printf("sudoku solved in %s", time.Since(start))
	log.Println("press CTRL+C to quit")

	log.Printf("\n%s\n", grid.ValuesToString())
	log.Printf("is grid correct: %v", grid.IsComplete())
	log.Println("Exit")
}
