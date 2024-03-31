package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/jtbonhomme/sugoku/internal/fonts"
	"github.com/jtbonhomme/sugoku/internal/sudoku"
	"github.com/jtbonhomme/sugoku/internal/text"
)

const (
	BoardX        float32 = 50
	BoardY        float32 = 50
	BoardWidth    float32 = 450
	BoardHeight   float32 = 450
	BoardCellSize float32 = 50
)

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.drawDebug(screen)

	g.drawFrame(screen)

	g.drawGridCandidates(screen)
	g.drawGridValues(screen)

	if g.grid.IsComplete() {
		g.drawFinish(screen)
	}

}

func (g *Game) drawFinish(screen *ebiten.Image) {
	text.DrawTextAtPos(
		screen, fonts.DefaultFont,
		5,
		35,
		"YOU WIN!",
		color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
	)
}

func (g *Game) drawGridCandidates(screen *ebiten.Image) {
	for row := 0; row < sudoku.Dim; row++ {
		for col := 0; col < sudoku.Dim; col++ {
			if !g.grid.CellIsEmpty(row, col) {
				continue
			}

			for i, c := range g.grid.Candidates(row, col) {
				y := i % 3
				x := (i - y) / 3
				if c != 0 {
					text.DrawTextAtPos(
						screen, fonts.SmallFont,
						55+row*50+y*14,
						65+col*50+x*14,
						fmt.Sprintf("%d", c),
						color.RGBA{R: 0xaf, G: 0xaf, B: 0xff, A: 0xaf},
					)
				}
			}
		}
	}
}

func (g *Game) drawGridValues(screen *ebiten.Image) {
	for row := 0; row < sudoku.Dim; row++ {
		for col := 0; col < sudoku.Dim; col++ {
			if g.grid.CellIsEmpty(row, col) {
				continue
			}

			text.DrawTextAtPos(
				screen, fonts.DefaultFont,
				65+row*50,
				85+col*50,
				fmt.Sprintf("%d", g.grid.Value(row, col)),
				color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
			)
		}
	}
}

func (g *Game) drawFrame(screen *ebiten.Image) {
	var x, y, width, height, strokeWidth, step float32

	x, y = BoardX, BoardY
	width, height = BoardWidth, BoardHeight
	strokeWidth = 1
	step = BoardCellSize

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

func (g *Game) drawDebug(screen *ebiten.Image) {
	//if !g.debug {
	//	return
	//}

	col, row := g.currentCase()
	if col != -1 && row != -1 {
		g.drawDebugCurrentCol(screen, col)
		g.drawDebugCurrentRow(screen, row)
		g.drawDebugCurrentBlock(screen, row, col)
	}
	debug := fmt.Sprintf("col, row = %d,   %d", col, row)
	text.DrawTextAtPos(
		screen, fonts.DefaultFont,
		5,
		35,
		debug,
		color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
	)
}

func (g *Game) drawDebugCurrentCol(screen *ebiten.Image, col int) {
	x := BoardX
	y := BoardX
	vector.DrawFilledRect(screen,
		x+float32(col)*BoardCellSize,
		y,
		BoardCellSize,
		BoardCellSize*float32(sudoku.Dim),
		color.RGBA{0x15, 0x15, 0x25, 0x30},
		true)
}

func (g *Game) drawDebugCurrentRow(screen *ebiten.Image, row int) {
	x := BoardX
	y := BoardX
	vector.DrawFilledRect(screen,
		x,
		y+float32(row)*BoardCellSize,
		BoardCellSize*float32(sudoku.Dim),
		BoardCellSize,
		color.RGBA{0x15, 0x15, 0x25, 0x30},
		true)
}

func (g *Game) drawDebugCurrentBlock(screen *ebiten.Image, row, col int) {
	x := BoardX
	y := BoardX
	row = row - row%3
	col = col - col%3
	vector.DrawFilledRect(screen,
		x+float32(col)*BoardCellSize,
		y+float32(row)*BoardCellSize,
		BoardCellSize*float32(sudoku.Dim/3),
		BoardCellSize*float32(sudoku.Dim/3),
		color.RGBA{0x25, 0x25, 0x35, 0x20},
		true)
}

func (g *Game) currentCase() (col, row int) {
	xCursor, yCursor := ebiten.CursorPosition()

	col = (xCursor - int(BoardX)) / int(BoardCellSize)
	row = (yCursor - int(BoardY)) / int(BoardCellSize)

	// cursor out of range
	if xCursor < int(BoardX) {
		col = -1
	}
	if xCursor > (int(BoardX) + int(BoardCellSize)*sudoku.Dim) {
		col = -1
	}

	if yCursor < int(BoardY) {
		row = -1
	}

	if yCursor > (int(BoardY) + int(BoardCellSize)*sudoku.Dim) {
		row = -1
	}

	return col, row
}
