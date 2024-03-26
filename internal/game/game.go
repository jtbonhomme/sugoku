package game

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

// Game manages all internal game mechanisms.
type Game struct {
	ScreenWidth     int
	ScreenHeight    int
	BackgroundColor color.Color
	grid            *sudoku.Grid
	wg              *sync.WaitGroup
}

// New creates a new game object.
func New(grid *sudoku.Grid, wg *sync.WaitGroup) *Game {
	g := &Game{
		ScreenWidth:     550,
		ScreenHeight:    550,
		BackgroundColor: color.RGBA{0x0b, 0x0d, 0x00, 0xff},
		wg:              wg,
		grid:            grid,
	}
	wg.Add(1)

	return g
}

// Run game loop.
func (g *Game) Run() error {

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Sudoko")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g.wg.Done()

	return ebiten.RunGame(g)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
