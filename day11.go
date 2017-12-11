package main

import (
	"log"
	"math"
	"strings"
)

func distanceFromStart(x int, y int, z int) int {
	dist := math.Max(math.Abs(float64(x)), math.Abs(float64(y)))
	dist = math.Max(dist, math.Abs(float64(z)))
	return int(dist)
}

func Day11(part int, data []byte) {
	steps := strings.Split(string(data), ",")
	log.Printf("Loaded path of %d steps", len(steps))
	current := struct {
		x int
		y int
		z int
	}{
		x: 0,
		y: 0,
		z: 0,
	}
	switch part {
	case 1:
		for _, step := range steps {
			switch step {
			case "n":
				current.y++
				current.z--
			case "ne":
				current.x++
				current.z--
			case "se":
				current.x++
				current.y--
			case "s":
				current.y--
				current.z++
			case "sw":
				current.x--
				current.z++
			case "nw":
				current.x--
				current.y++
			}
		}
		log.Printf("Child position is (%d,%d,%d) which is %d steps away", current.x, current.y, current.z, distanceFromStart(current.x, current.y, current.z))
	case 2:
		maxDist := 0
		for _, step := range steps {
			switch step {
			case "n":
				current.y++
				current.z--
			case "ne":
				current.x++
				current.z--
			case "se":
				current.x++
				current.y--
			case "s":
				current.y--
				current.z++
			case "sw":
				current.x--
				current.z++
			case "nw":
				current.x--
				current.y++
			}
			dist := distanceFromStart(current.x, current.y, current.z)
			if dist > maxDist {
				maxDist = dist
			}
		}
		log.Printf("Maximum distance was %d steps away", maxDist)
	}
}
