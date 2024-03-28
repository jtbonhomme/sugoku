package sudoku

import (
	"testing"
)

var testGridOK = Grid{
	values: [Dim][Dim]byte{
		{9, 4, 8, 1, 5, 3, 2, 6, 7},
		{2, 6, 1, 4, 9, 7, 8, 5, 3},
		{3, 5, 7, 2, 6, 8, 4, 1, 9},

		{1, 8, 5, 9, 7, 4, 3, 2, 6},
		{6, 9, 4, 8, 3, 2, 5, 7, 1},
		{7, 3, 2, 5, 1, 6, 9, 8, 4},

		{4, 2, 6, 3, 8, 1, 7, 9, 5},
		{8, 7, 9, 6, 4, 5, 1, 3, 2},
		{5, 1, 3, 7, 2, 9, 6, 4, 8},
	},
}

var testGridKO = Grid{
	values: [Dim][Dim]byte{
		{4, 9, 9, 1, 5, 3, 2, 6, 7},
		{2, 6, 1, 4, 9, 7, 8, 5, 3},
		{3, 5, 7, 2, 6, 8, 4, 1, 9},

		{1, 8, 5, 9, 7, 4, 3, 2, 6},
		{6, 9, 4, 8, 3, 2, 5, 7, 1},
		{7, 3, 2, 5, 1, 6, 9, 8, 4},

		{4, 2, 6, 3, 8, 1, 7, 9, 5},
		{8, 7, 9, 6, 4, 5, 1, 3, 2},
		{5, 1, 3, 7, 2, 9, 6, 4, 8},
	},
}

func TestIsCompleteKO(t *testing.T) {
	if testGridKO.IsRowComplete(0) {
		t.Errorf("expected row 0 not to be complete but is actually")
	}

	if !testGridKO.IsRowComplete(1) {
		t.Errorf("expected row 1 to be complete but is not")
	}

	if testGridKO.IsColComplete(0) {
		t.Errorf("expected col 0 not to be complete but is actually")
	}

	if !testGridKO.IsColComplete(3) {
		t.Errorf("expected col 3 to be complete but is not")
	}

	if testGridKO.IsBlockComplete(0, 0) {
		t.Errorf("expected block (0, 0) not to be complete but is actually")
	}

	if !testGridKO.IsBlockComplete(3, 0) {
		t.Errorf("expected block (3, 0) to be complete but is not")
	}
}

func TestIsRowComplete(t *testing.T) {
	for row := 0; row < Dim; row++ {
		if !testGridOK.IsRowComplete(row) {
			t.Errorf("expected row %d to be complete but is not", row)
		}
	}
}

func TestIsColComplete(t *testing.T) {
	for col := 0; col < Dim; col++ {
		if !testGridOK.IsRowComplete(col) {
			t.Errorf("expected col %d to be complete but is not", col)
		}
	}
}

func TestIsBlockComplete(t *testing.T) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if !testGridOK.IsBlockComplete(row*3, col*3) {
				t.Errorf("expected block (%d,%d) to be complete but is not", row, col)
			}
		}
	}
}
