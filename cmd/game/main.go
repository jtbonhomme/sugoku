package main

import (
	"flag"
	"log"
	"sync"

	"github.com/jtbonhomme/sugoku/internal/game"
	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func main() {
	var err error
	var filename string

	flag.StringVar(&filename, "f", "_examples/not-unique-solution.json", "initial grid (default is _examples/backtracking.json)")
	flag.Parse()

	grid := sudoku.NewFromFile(filename)

	var wg sync.WaitGroup

	g := game.New(grid, &wg)

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exit")
}
