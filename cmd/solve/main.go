package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/jtbonhomme/sugoku/internal/game"
	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

func main() {
	var err error
	var filename string
	flag.StringVar(&filename, "f", "_examples/backtracking.json", "initial grid (default is _examples/backtracking.json)")
	flag.Parse()

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := sudoku.Grid{}
	err = json.Unmarshal(fileBytes, &grid)
	if err != nil {
		log.Fatal(err)
	}

	g := game.New(&grid)

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Exit")
}
