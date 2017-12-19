package main

import (
	"log"
	"strings"
)

const DIRECTION_UP = 0
const DIRECTION_RIGHT = 1
const DIRECTION_DOWN = 2
const DIRECTION_LEFT = 3

type Diagram struct {
	diag      [][]byte
	x         int
	y         int
	height    int
	width     int
	path      []byte
	direction int
	counter   int
}

func NewDiagram(data []byte) *Diagram {
	lines := strings.Split(string(data), "\n")
	diag := make([][]byte, len(lines))
	for i, line := range lines {
		diag[i] = []byte(line)
	}
	result := &Diagram{
		diag:      diag,
		x:         0,
		y:         0,
		height:    len(diag),
		width:     len(diag[0]),
		path:      nil,
		direction: DIRECTION_DOWN,
		counter:   1,
	}
	for i := 0; i < len(diag[0]); i++ {
		if diag[0][i] == '|' {
			result.x = i
			break
		}
	}
	return result
}

func (d *Diagram) Width() int {
	return d.width
}

func (d *Diagram) Height() int {
	return d.height
}

func (d *Diagram) Empty(x int, y int) bool {
	if y < 0 || y >= d.height || x < 0 || x >= d.width {
		return true
	}
	return d.diag[y][x] == ' '
}

func (d *Diagram) RecordPath() {
	if d.current() != '|' && d.current() != '-' && d.current() != '+' {
		d.path = append(d.path, d.current())
	}
}

func (d *Diagram) current() byte {
	return d.diag[d.y][d.x]
}

func (d *Diagram) chooseDirection() bool {
	if d.direction != DIRECTION_UP && !d.Empty(d.x, d.y+1) {
		d.direction = DIRECTION_DOWN
		return true
	}
	if d.direction != DIRECTION_LEFT && !d.Empty(d.x+1, d.y) {
		d.direction = DIRECTION_RIGHT
		return true
	}
	if d.direction != DIRECTION_DOWN && !d.Empty(d.x, d.y-1) {
		d.direction = DIRECTION_UP
		return true
	}
	if d.direction != DIRECTION_RIGHT && !d.Empty(d.x-1, d.y) {
		d.direction = DIRECTION_LEFT
		return true
	}
	return false
}

func (d *Diagram) move() bool {
	switch d.direction {
	case DIRECTION_DOWN:
		if d.Empty(d.x, d.y+1) {
			return false
		}
		d.y++
		return true
	case DIRECTION_UP:
		if d.Empty(d.x, d.y-1) {
			return false
		}
		d.y--
		return true
	case DIRECTION_RIGHT:
		if d.Empty(d.x+1, d.y) {
			return false
		}
		d.x++
		return true
	case DIRECTION_LEFT:
		if d.Empty(d.x-1, d.y) {
			return false
		}
		d.x--
		return true
	}
	return false
}

func (d *Diagram) Next() bool {
	if d.move() {
		d.counter++
		return true
	} else if d.chooseDirection() {
		return d.Next()
	}
	return false
}

func (d *Diagram) FindExit() {
	for d.Next() {
		d.RecordPath()
	}
}

func (d *Diagram) Path() string {
	return string(d.path)
}

func (d *Diagram) Counter() int {
	return d.counter
}

func Day19(part int, data []byte) {
	diag := NewDiagram(data)
	log.Printf("Loaded diagram size %dx%d", diag.Width(), diag.Height())
	diag.FindExit()
	switch part {
	case 1:
		log.Printf("Reached end of path, letters seen %s", diag.Path())
	case 2:
		log.Printf("Reached end of path in %d steps", diag.Counter())
	}
}
