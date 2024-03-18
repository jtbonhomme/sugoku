package game

import (
	"time"
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

func (g *Game) IsValid(position int) bool {
	time.Sleep(time.Millisecond * 7)
	// Si on est à la 82e case (on sort du tableau)
	if position == 9*9 {
		return true
	}

	// On récupère les coordonnées de la case
	i := position / 9
	j := position % 9

	// Si la case n'est pas vide, on passe à la suivante (appel récursif)
	if g.grid[i][j] != 0 {
		return g.IsValid(position + 1)
	}

	// Backtracking
	for k := 1; k <= 9; k++ {
		if g.grid.MissingInRow(byte(k), i) == true &&
			g.grid.MissingInColumn(byte(k), j) == true &&
			g.grid.MissingInBlock(byte(k), i, j) == true {
			g.grid[i][j] = byte(k)

			if g.IsValid(position+1) == true {
				return true
			}
		}
	}
	g.grid[i][j] = 0

	return false
}
