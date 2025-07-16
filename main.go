package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
	Invalid Direction = -1
)

var directions = []string{"N", "E", "S", "W"}

func (d Direction) String() string {
	return directions[d]
}

type Position struct {
	x   int
	y   int
	dir Direction
}

type Plateau struct {
	width  int
	height int
}

func (p *Position) rotateLeft() {
	p.dir = (p.dir + 3) % 4
}

func (p *Position) rotateRight() {
	p.dir = (p.dir + 1) % 4
}

func (p *Position) moveForward(plateau Plateau) {
	switch p.dir {
	case North:
		if p.y < plateau.height {
			p.y++
		}
	case East:
		if p.x < plateau.width {
			p.x++
		}
	case South:
		if p.y > 0 {
			p.y--
		}
	case West:
		if p.x > 0 {
			p.x--
		}
	}
}

func parseDirection(s string) (Direction, error) {
	switch s {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	default:
		return Invalid, fmt.Errorf("invalid direction: %s", s)

	}
}

func parsePosition(line string, plateau Plateau) (Position, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return Position{}, fmt.Errorf("invalid position line: %s", line)
	}
	x, err1 := strconv.Atoi(parts[0])
	y, err2 := strconv.Atoi(parts[1])
	dir, err3 := parseDirection(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return Position{}, fmt.Errorf("invalid position values: %s", line)
	}
	if x < 0 || x > plateau.width || y < 0 || y > plateau.height {
		return Position{}, fmt.Errorf("position out of plateau bounds: %s", line)
	}
	return Position{x, y, dir}, nil
}

func validateInstructions(instr string) error {
	for _, c := range instr {
		if c != 'L' && c != 'R' && c != 'M' {
			return fmt.Errorf("invalid instruction: %c", c)
		}
	}
	return nil
}

func parsePlateau(line string) (Plateau, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return Plateau{}, fmt.Errorf("invalid plateau line: %s", line)
	}
	width, err1 := strconv.Atoi(parts[0])
	height, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || width < 0 || height < 0 {
		return Plateau{}, fmt.Errorf("invalid plateau size: %s", line)

	}
	return Plateau{width, height}, nil
}

func readPlateau(scanner *bufio.Scanner) (Plateau, error) {
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		return parsePlateau(text)
	}
	return Plateau{}, fmt.Errorf("missing plateau line")
}

func readRoverPair(scanner *bufio.Scanner) (string, string, error) {
	posLine := ""
	for posLine == "" && scanner.Scan() {
		posLine = strings.TrimSpace(scanner.Text())
	}
	if posLine == "" {
		return "", "", fmt.Errorf("no more rover positions")
	}

	instrLine := ""
	for instrLine == "" && scanner.Scan() {
		instrLine = strings.TrimSpace(scanner.Text())
	}
	if instrLine == "" {
		return "", "", fmt.Errorf("missing instructions for rover after: %s", posLine)
	}
	return posLine, instrLine, nil
}

func processRover(posLine, instrLine string, plateau Plateau) error {
	pos, err := parsePosition(posLine, plateau)
	if err != nil {
		return fmt.Errorf("rover position error: %v", err)
	}
	if err := validateInstructions(instrLine); err != nil {
		return fmt.Errorf("instruction error: %v", err)
	}
	for _, cmd := range instrLine {
		switch cmd {
		case 'L':
			pos.rotateLeft()
		case 'R':
			pos.rotateRight()
		case 'M':
			pos.moveForward(plateau)
		}
	}
	fmt.Printf("%d %d %s\n", pos.x, pos.y, pos.dir.String())
	return nil
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) > 1 && os.Args[1] != "" {
		inputFile = os.Args[1]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not open file '%s'.\nUsage: %s [inputfile.txt]\n", inputFile, os.Args[0])
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plateau, err := readPlateau(scanner)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Plateau error: %v\n", err)
		os.Exit(1)
	}

	for {
		posLine, instrLine, err := readRoverPair(scanner)
		if err != nil {
			if err.Error() != "no more rover positions" {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			break
		}
		if err := processRover(posLine, instrLine, plateau); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		os.Exit(1)
	}
}
