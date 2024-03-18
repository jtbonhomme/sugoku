package sudoku

const DIM int = 9

type Grid [DIM][DIM]byte

func (g *Grid) MissingInRow(k byte, i int) bool {
	for j := 0; j < 9; j++ {
		if g[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) MissingInColumn(k byte, j int) bool {
	for i := 0; i < 9; i++ {
		if g[i][j] == k {
			return false
		}
	}

	return true
}

func (g *Grid) MissingInBlock(k byte, i int, j int) bool {
	_i := i - (i % 3)
	_j := j - (j % 3)
	for i := _i; i < _i+3; i++ {
		for j := _j; j < _j+3; j++ {
			if g[i][j] == k {
				return false
			}
		}
	}

	return true
}
