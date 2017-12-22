package main

import (
	"log"
	"strings"
)

const SPORIFICA_BURSTS = 10000
const SPORIFICA_BURSTS2 = 10000000

const SPORIFICA_CLEAN = 0
const SPORIFICA_WEAKENED = 1
const SPORIFICA_INFECTED = 2
const SPORIFICA_FLAGGED = 3

type Sporifica struct {
	pos       Point
	direction int
}

func NewSporifica(x int, y int) *Sporifica {
	return &Sporifica{
		pos:       Point{x: x, y: y},
		direction: DIRECTION_UP,
	}
}

func (s *Sporifica) Move() {
	switch s.direction {
	case DIRECTION_UP:
		s.pos.y--
	case DIRECTION_DOWN:
		s.pos.y++
	case DIRECTION_LEFT:
		s.pos.x--
	case DIRECTION_RIGHT:
		s.pos.x++
	default:
		log.Fatalf("Unknown direction %d", s.direction)
	}
}

func (s *Sporifica) UpdateDirection(status int) {
	switch status {
	case SPORIFICA_CLEAN:
		s.direction--
	case SPORIFICA_WEAKENED:
	case SPORIFICA_INFECTED:
		s.direction++
	case SPORIFICA_FLAGGED:
		s.direction += 2
	}
	s.direction %= 4
	if s.direction < 0 {
		s.direction += 4
	}
}

func Day22(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	grid := make(map[Point]int)
	for i, line := range lines {
		for j, v := range line {
			if v == '#' {
				point := Point{x: j, y: i}
				grid[point] = SPORIFICA_INFECTED
			}
		}
	}
	initialSize := len(lines)
	sporifica := NewSporifica(initialSize/2, initialSize/2)
	log.Printf("Loaded grid of size %d with %d infected cells", initialSize, len(grid))
	switch part {
	case 1:
		infections := 0
		for i := 0; i < SPORIFICA_BURSTS; i++ {
			sporifica.UpdateDirection(grid[sporifica.pos])
			if grid[sporifica.pos] == SPORIFICA_INFECTED {
				delete(grid, sporifica.pos)
			} else {
				grid[sporifica.pos] = SPORIFICA_INFECTED
				infections++
			}
			sporifica.Move()
		}
		log.Printf("After %d bursts there are %d infection bursts", SPORIFICA_BURSTS, infections)
	case 2:
		infections := 0
		for i := 0; i < SPORIFICA_BURSTS2; i++ {
			sporifica.UpdateDirection(grid[sporifica.pos])
			switch grid[sporifica.pos] {
			case SPORIFICA_CLEAN:
				grid[sporifica.pos] = SPORIFICA_WEAKENED
			case SPORIFICA_WEAKENED:
				grid[sporifica.pos] = SPORIFICA_INFECTED
				infections++
			case SPORIFICA_INFECTED:
				grid[sporifica.pos] = SPORIFICA_FLAGGED
			case SPORIFICA_FLAGGED:
				delete(grid, sporifica.pos)
			}
			sporifica.Move()
		}
		log.Printf("After %d bursts there are %d infection bursts", SPORIFICA_BURSTS2, infections)
	}
}
