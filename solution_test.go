package square

import (
	"testing"
)

func TestSquare(t *testing.T) {
	expectedSize := 6.25
	squareSize := CalcSquare(2.5, SidesSquare)
	if squareSize != expectedSize {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expectedSize, squareSize)
	}
}

func TestTriangle(t *testing.T) {
	expectedSize := 8.768507213317442
	squareSize := CalcSquare(4.5, SidesTriangle)
	if squareSize != expectedSize {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expectedSize, squareSize)
	}
}

func TestHexagon(t *testing.T) {
	expectedSize := 0.0
	var SidesHexagon intCustomType = 6
	squareSize := CalcSquare(2.5, SidesHexagon)
	if squareSize != expectedSize {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expectedSize, squareSize)
	}
}

func TestCircle(t *testing.T) {
	expectedSize := 7.0685834705770345
	squareSize := CalcSquare(1.5, SidesCircle)
	if squareSize != expectedSize {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expectedSize, squareSize)
	}
}
