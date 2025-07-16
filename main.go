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
	Invalid Direction = -1 // Explicit invalid value for Direction
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
	p.dir = (p.dir + 3) % 4 // equivalent to -1 mod 4 tes
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

func parseDirection(s string) Direction {
	switch s {
	case "N":
		return North
	case "E":
		return East
	case "S":
		return South
	case "W":
		return West
	default:
		fmt.Fprintf(os.Stderr, "invalid direction: %s\n", s)
		return Invalid // Use -1 as an explicit invalid value for Direction

	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error openning file: %s\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	var plateau Plateau

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		if line == 0 {
			parts := strings.Split(text, " ")
			plateau.width, _ = strconv.Atoi(parts[0])
			plateau.height, _ = strconv.Atoi(parts[1])
			line++
			continue
		}

		// Rover position
		parts := strings.Split(text, " ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		dir := parseDirection(parts[2])
		pos := Position{x, y, dir}

		// Next line is instructions
		if !scanner.Scan() {
			break
		}
		instructions := scanner.Text()
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

		fmt.Printf("%d %d %s\n", pos.x, pos.y, pos.dir.String())
		// Move to the next rover

		line++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)

	}
}
