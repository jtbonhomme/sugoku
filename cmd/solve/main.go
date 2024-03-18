package main

import (
	"fmt"
	//"os"
	//"os/signal"
	//"syscall"
	//"time"

	"github.com/jtbonhomme/sugoku/internal/game"
)

func main() {
	g := game.New()

	fmt.Println("Start game")
	err := g.Run()
	if err != nil {
		panic(err)
	}
	/*
	   grid.Display()
	   grid.IsValid(0)
	   grid.Display()

	   // Waiting signal
	   intChan := make(chan os.Signal, 1)
	   signal.Notify(intChan, os.Interrupt, syscall.SIGTERM)

	   <-intChan
	*/
	fmt.Println("Exit")
}
