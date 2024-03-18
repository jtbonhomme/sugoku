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

	fmt.Println("Exit")
}
