package main

import (
	"log"
	"strconv"
	"strings"
)

func Day5(part int, data []byte) {
	var err error
	lines := strings.Split(string(data), "\n")
	maze := make([]int, len(lines))
	for i, line := range lines {
		maze[i], err = strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed parsing number %s at line %d", line, i)
		}
	}
	log.Printf("Loaded maze of %d numbers", len(maze))
	switch part {
	case 1:
		current := 0
		counter := 0
		for current >= 0 && current < len(maze) {
			counter++
			maze[current]++
			current += maze[current] - 1
		}
		log.Printf("Jumped out of maze after %d jumps", counter)
	case 2:
		current := 0
		counter := 0
		for current >= 0 && current < len(maze) {
			counter++
			offset := maze[current]
			if maze[current] >= 3 {
				maze[current]--
			} else {
				maze[current]++
			}
			current += offset
		}
		log.Printf("Jumped out of maze after %d jumps", counter)
	}
}
