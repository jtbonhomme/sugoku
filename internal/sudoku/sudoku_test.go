package sudoku

import (
	"testing"
)

var testGridOK = Grid{
	values: [Dim][Dim]byte{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 3, 2, 1, 6, 5, 4},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},

		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{7, 9, 8, 1, 3, 2, 4, 6, 5},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},

		{3, 1, 5, 2, 7, 4, 9, 6, 8},
		{8, 7, 9, 2, 6, 3, 4, 1, 5},
		{5, 6, 4, 8, 1, 7, 3, 9, 2},
	},
}

func TestIsRowComplete(t *testing.T) {
	for row := 0; row < Dim; row++ {
		if !testGridOK.IsRowComplete(row) {
			t.Errorf("expected row %d to be complete but is not", row)
		}
	}
}
