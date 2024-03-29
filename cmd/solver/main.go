package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/jtbonhomme/sugoku/internal/game"
	"github.com/jtbonhomme/sugoku/internal/solver"
	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func main() {
	var err error
	var filename string
	var speed int
	var noCandidates bool

	flag.StringVar(&filename, "f", "_examples/not-unique-solution.json", "initial grid (default is _examples/backtracking.json)")
	flag.IntVar(&speed, "s", 1, "speed resolution (default is 1)")
	flag.BoolVar(&noCandidates, "n", false, "do not display candidates (default is false)")
	flag.Parse()

	grid := sudoku.NewFromFile(filename)

	var wg sync.WaitGroup
	var m sync.Mutex
	m.Lock()

	g := game.New(grid, &wg)

	go func() {
		wg.Wait()
		start := time.Now()
		log.Println("start backtracking sudoku")
		solver.SolveWithBacktracking(0, grid, speed, noCandidates)
		log.Printf("sudoku solved in %s", time.Since(start))
		log.Println("press CTRL+C to quit")
	}()

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exit")
}
