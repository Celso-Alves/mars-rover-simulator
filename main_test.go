package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGivenNorthDirectionWhenRotateLeftThenFaceWest(t *testing.T) {
	// Arrange
	pos := Position{x: 0, y: 0, dir: North}

	// Act
	pos.rotateLeft()

	// Assert
	if pos.dir != West {
		t.Errorf("expected West, got %s", pos.dir.String())
	}
}

func TestGivenEastDirectionWhenRotateRightThenFaceSouth(t *testing.T) {
	// Arrange
	pos := Position{x: 0, y: 0, dir: East}

	// Act
	pos.rotateRight()

	// Assert
	if pos.dir != South {
		t.Errorf("expected South, got %s", pos.dir.String())
	}
}

func TestGivenNorthDirectionWhenMoveForwardThenIncreaseY(t *testing.T) {
	// Arrange
	pos := Position{x: 1, y: 1, dir: North}
	plateau := Plateau{width: 5, height: 5}

	// Act
	pos.moveForward(plateau)

	// Assert
	if pos.x != 1 || pos.y != 2 {
		t.Errorf("expected position (1,2), got (%d,%d)", pos.x, pos.y)
	}
}

func TestGivenWestDirectionWhenMoveForwardOnEdgeThenDoNotMove(t *testing.T) {
	// Arrange
	pos := Position{x: 0, y: 0, dir: West}
	plateau := Plateau{width: 5, height: 5}

	// Act
	pos.moveForward(plateau)

	// Assert
	if pos.x != 0 || pos.y != 0 {
		t.Errorf("expected position (0,0), got (%d,%d)", pos.x, pos.y)
	}
}

func TestGivenInstructionsWhenExecuteAllThenReachFinalPosition(t *testing.T) {
	// Arrange
	pos := Position{x: 1, y: 2, dir: North}
	plateau := Plateau{width: 5, height: 5}
	instructions := "LMLMLMLMM"

	// Act
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			pos.rotateLeft()
		case 'R':
			pos.rotateRight()
		case 'M':
			pos.moveForward(plateau)
		}
	}

	// Assert
	if pos.x != 1 || pos.y != 3 || pos.dir != North {
		t.Errorf("expected position 1 3 N, got %d %d %s", pos.x, pos.y, pos.dir.String())
	}
}

func TestGivenValidInputFileWhenRunMainThenPrintExpectedOutput(t *testing.T) {
	// Arrange
	input := `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM`
	tempFile := "input.txt"
	err := os.WriteFile(tempFile, []byte(input), 0644)
	if err != nil {
		t.Fatalf("failed to write input file: %v", err)
	}
	defer os.Remove(tempFile)

	// Redirect stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	main()
	w.Close()
	os.Stdout = oldStdout

	// Read output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Assert
	expected := "1 3 N\n5 1 E\n"
	if output != expected {
		t.Errorf("unexpected output.\nGot:\n%s\nWant:\n%s", output, expected)
	}
}
func TestParseDirectionValidInputs(t *testing.T) {
	tests := []struct {
		input    string
		expected Direction
	}{
		{"N", North},
		{"E", East},
		{"S", South},
		{"W", West},
	}

	for _, tt := range tests {
		got := parseDirection(tt.input)
		if got != tt.expected {
			t.Errorf("parseDirection(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestParseDirectionInvalidInputPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("parseDirection did not panic on invalid input")
		}
	}()
	parseDirection("X")
}
