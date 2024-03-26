package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/jtbonhomme/sugoku/internal/sudoku"
)

// Game manages all internal game mechanisms.
type Game struct {
	ScreenWidth     int
	ScreenHeight    int
	BackgroundColor color.Color
	grid            *sudoku.Grid
}

// New creates a new game object.
func New(grid *sudoku.Grid) *Game {
	g := &Game{
		ScreenWidth:     800,
		ScreenHeight:    600,
		BackgroundColor: color.RGBA{0x0b, 0x0d, 0x00, 0xff},
	}

	g.grid = grid

	return g
}

// Run game loop.
func (g *Game) Run() error {

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Sudoko")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	go func() {
		g.IsValid(0)
	}()

	return ebiten.RunGame(g)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
