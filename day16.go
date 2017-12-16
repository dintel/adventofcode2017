package main

import (
	"log"
	"strconv"
	"strings"
)

const DANCE_PROGRAMS = 16
const DANCE_REPETITIONS = 1000000000

type ProgramDance struct {
	size      int
	positions []string
	names     map[string]int
}

func NewProgramDance(num int) *ProgramDance {
	result := &ProgramDance{
		size:      num,
		positions: make([]string, num),
		names:     make(map[string]int),
	}
	for i := 0; i < num; i++ {
		name := string('a' + i)
		result.positions[i] = name
		result.names[name] = i
	}
	return result
}

func (dance *ProgramDance) Spin(x int) {
	for name := range dance.names {
		dance.names[name] = (dance.names[name] + x) % dance.size
		dance.positions[dance.names[name]] = name
	}
}

func (dance *ProgramDance) Exchange(first int, second int) {
	dance.positions[first], dance.positions[second] = dance.positions[second], dance.positions[first]
	dance.names[dance.positions[first]] = first
	dance.names[dance.positions[second]] = second
}

func (dance *ProgramDance) Partner(first string, second string) {
	dance.names[first], dance.names[second] = dance.names[second], dance.names[first]
	dance.positions[dance.names[first]] = first
	dance.positions[dance.names[second]] = second
}

func (dance *ProgramDance) Standing() string {
	result := ""
	for _, program := range dance.positions {
		result += program
	}
	return result
}

func (dance *ProgramDance) Dance(steps []string) string {
	for i, step := range steps {
		move := step[0:1]
		params := strings.Split(step[1:], "/")
		switch move {
		case "s":
			offset, err := strconv.Atoi(params[0])
			if err != nil {
				log.Fatalf("Failed parsing spin offset '%s' at move %d", params[0], i)
			}
			dance.Spin(offset)
		case "x":
			first, err := strconv.Atoi(params[0])
			if err != nil {
				log.Fatalf("Failed parsing first exchange parameter '%s' at move %d", params[0], i)
			}
			second, err := strconv.Atoi(params[1])
			if err != nil {
				log.Fatalf("Failed parsing second exchange parameter '%s' at move %d", params[1], i)
			}
			dance.Exchange(first, second)
		case "p":
			dance.Partner(params[0], params[1])
		default:
			log.Fatalf("Unknown move %s at step %d", move, i)
		}
	}
	return dance.Standing()
}

func Day16(part int, data []byte) {
	steps := strings.Split(string(data), ",")
	programDance := NewProgramDance(16)
	log.Printf("Loaded %d dance steps", len(steps))
	switch part {
	case 1:
		programDance.Dance(steps)
		log.Printf("Final stading is %s", programDance.Standing())
	case 2:
		results := make(map[string]int)
		iterations := make(map[int]string)
		offset := 0
		delta := 0
		for i := 0; i < DANCE_REPETITIONS; i++ {
			current := i + 1
			standing := programDance.Dance(steps)
			if previous, exists := results[standing]; exists {
				log.Printf("Found repetition after dance iteration %d (same as on iteration %d)", current, previous)
				offset = previous
				delta = current - previous
				break
			}
			results[standing] = current
			iterations[current] = standing

		}
		final := (DANCE_REPETITIONS - offset) % delta
		log.Printf("Final stading is %s", iterations[offset+final])
	}
}
