package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
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

	flag.StringVar(&filename, "f", "_examples/backtracking.json", "initial grid (default is _examples/backtracking.json)")
	flag.IntVar(&speed, "s", 1, "speed resolution (default is 1)")
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

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)

	var wg sync.WaitGroup
	var m sync.Mutex
	m.Lock()

	g := game.New(&grid, &wg)

	go func() {
		wg.Wait()
		start := time.Now()
		log.Println("start backtracking sudoku")
		solver.SolveWithBacktracking(0, &grid, speed)
		log.Printf("sudoku solved in %s", time.Since(start))
		log.Println("press CTRL+C to quit")
	}()

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Block until a signal is received.
	//s := <-c
	//log.Println("Got signal:", s)
	log.Println("Exit")
}
