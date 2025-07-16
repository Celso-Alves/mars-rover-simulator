package main

import (
	"bufio"
	"strings"
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

func TestGivenValidDirectionStringWhenParseDirectionThenReturnCorrectDirection(t *testing.T) {
	// Arrange
	tests := []struct {
		input    string
		expected Direction
	}{
		{"N", North},
		{"E", East},
		{"S", South},
		{"W", West},
	}
	// Act & Assert
	for _, tt := range tests {
		got, err := parseDirection(tt.input)
		if err != nil {
			t.Errorf("parseDirection(%q) returned error: %v", tt.input, err)
		}
		if got != tt.expected {
			t.Errorf("parseDirection(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestGivenInvalidDirectionStringWhenParseDirectionThenReturnError(t *testing.T) {
	// Arrange
	// Act
	_, err := parseDirection("X")
	// Assert
	if err == nil {
		t.Errorf("parseDirection did not return error on invalid input")
	}
}

func TestGivenDirectionWhenCallStringThenReturnCorrectString(t *testing.T) {
	// Arrange
	tests := []struct {
		dir      Direction
		expected string
	}{
		{North, "N"},
		{East, "E"},
		{South, "S"},
		{West, "W"},
	}
	// Act & Assert
	for _, tt := range tests {
		if tt.dir.String() != tt.expected {
			t.Errorf("Direction.String() = %s, want %s", tt.dir.String(), tt.expected)
		}
	}
}

func TestGivenValidPositionStringWhenParsePositionThenReturnCorrectPosition(t *testing.T) {
	// Arrange
	plateau := Plateau{width: 5, height: 5}
	// Act
	pos, err := parsePosition("1 2 N", plateau)
	// Assert
	if err != nil {
		t.Errorf("parsePosition returned error: %v", err)
	}
	if pos.x != 1 || pos.y != 2 || pos.dir != North {
		t.Errorf("parsePosition returned wrong position: %+v", pos)
	}
}

func TestGivenInvalidPositionStringsWhenParsePositionThenReturnError(t *testing.T) {
	// Arrange
	plateau := Plateau{width: 5, height: 5}
	// Act & Assert
	_, err := parsePosition("1 2", plateau)
	if err == nil {
		t.Error("expected error for missing direction")
	}
	_, err = parsePosition("a b N", plateau)
	if err == nil {
		t.Error("expected error for non-integer coordinates")
	}
	_, err = parsePosition("10 10 N", plateau)
	if err == nil {
		t.Error("expected error for out of bounds")
	}
	_, err = parsePosition("1 2 X", plateau)
	if err == nil {
		t.Error("expected error for invalid direction")
	}
}

func TestGivenValidInstructionsWhenValidateInstructionsThenReturnNil(t *testing.T) {
	// Arrange
	// Act
	err := validateInstructions("LRMLMRM")
	// Assert
	if err != nil {
		t.Errorf("validateInstructions returned error for valid input: %v", err)
	}
}

func TestGivenInvalidInstructionsWhenValidateInstructionsThenReturnError(t *testing.T) {
	// Arrange
	// Act
	err := validateInstructions("LRMX")
	// Assert
	if err == nil {
		t.Error("expected error for invalid instruction")
	}
}

func TestGivenValidPlateauStringWhenParsePlateauThenReturnCorrectPlateau(t *testing.T) {
	// Arrange
	// Act
	plat, err := parsePlateau("5 5")
	// Assert
	if err != nil {
		t.Errorf("parsePlateau returned error: %v", err)
	}
	if plat.width != 5 || plat.height != 5 {
		t.Errorf("parsePlateau returned wrong plateau: %+v", plat)
	}
}

func TestGivenInvalidPlateauStringsWhenParsePlateauThenReturnError(t *testing.T) {
	// Arrange
	// Act & Assert
	_, err := parsePlateau("5")
	if err == nil {
		t.Error("expected error for missing height")
	}
	_, err = parsePlateau("a b")
	if err == nil {
		t.Error("expected error for non-integer values")
	}
	_, err = parsePlateau("-1 5")
	if err == nil {
		t.Error("expected error for negative width")
	}
}

func TestGivenPlateauInputWhenReadPlateauThenReturnPlateau(t *testing.T) {
	// Arrange
	input := "5 5\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Act
	plat, err := readPlateau(scanner)
	// Assert
	if err != nil {
		t.Errorf("readPlateau returned error: %v", err)
	}
	if plat.width != 5 || plat.height != 5 {
		t.Errorf("readPlateau returned wrong plateau: %+v", plat)
	}
}

func TestGivenEmptyInputWhenReadPlateauThenReturnError(t *testing.T) {
	// Arrange
	scanner := bufio.NewScanner(strings.NewReader("\n"))
	// Act
	_, err := readPlateau(scanner)
	// Assert
	if err == nil {
		t.Error("expected error for missing plateau line")
	}
}

func TestGivenValidRoverPairInputWhenReadRoverPairThenReturnPositionAndInstructions(t *testing.T) {
	// Arrange
	input := "1 2 N\nLMLMLMLMM\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Act
	posLine, instrLine, err := readRoverPair(scanner)
	// Assert
	if err != nil {
		t.Errorf("readRoverPair returned error: %v", err)
	}
	if posLine != "1 2 N" || instrLine != "LMLMLMLMM" {
		t.Errorf("readRoverPair returned wrong lines: %s, %s", posLine, instrLine)
	}
}

func TestGivenMissingInstructionsWhenReadRoverPairThenReturnError(t *testing.T) {
	// Arrange
	input := "1 2 N\n\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Act
	_, _, err := readRoverPair(scanner)
	// Assert
	if err == nil {
		t.Error("expected error for missing instructions")
	}
}

func TestGivenEmptyInputWhenReadRoverPairThenReturnError(t *testing.T) {
	// Arrange
	scanner := bufio.NewScanner(strings.NewReader(""))
	// Act
	_, _, err := readRoverPair(scanner)
	// Assert
	if err == nil {
		t.Error("expected error for no more rover positions")
	}
}
