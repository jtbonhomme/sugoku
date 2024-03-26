package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/jtbonhomme/sugoku/internal/fonts"
	"github.com/jtbonhomme/sugoku/internal/text"
)

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.drawFrame(screen)
	g.drawGrid(screen)
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if g.grid.Values[row][col] == 0 {
				continue
			}

			text.DrawTextAtPos(
				screen, fonts.DefaultFont,
				65+row*50,
				85+col*50,
				fmt.Sprintf("%d", g.grid.Values[row][col]),
				color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
			)
		}
	}
}

func (g *Game) drawFrame(screen *ebiten.Image) {
	var x, y, width, height, strokeWidth, step float32

	x, y = 50, 50
	width, height = 450, 450
	strokeWidth = 1
	step = 50

	for i := 0; i < 10; i++ {
		vector.StrokeLine(screen,
			x, y+step*float32(i), x+width, y+step*float32(i),
			strokeWidth, color.RGBA{0x8b, 0x8d, 0x80, 0xff}, false)
		vector.StrokeLine(screen,
			x+step*float32(i), y, x+step*float32(i), y+height,
			strokeWidth, color.RGBA{0x8b, 0x8d, 0x80, 0xff}, false)
	}

	for i := 0; i < 4; i++ {
		vector.StrokeLine(screen,
			x, y+step*float32(3*i), x+width, y+step*float32(3*i),
			3*strokeWidth, color.RGBA{0xff, 0xff, 0xff, 0xff}, false)
		vector.StrokeLine(screen,
			x+step*float32(3*i), y, x+step*float32(3*i), y+height,
			3*strokeWidth, color.RGBA{0xff, 0xff, 0xff, 0xff}, false)
	}
}
