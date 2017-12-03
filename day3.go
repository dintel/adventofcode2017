package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Day3(part int, data []byte) {
	num, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		log.Fatalf("Failed parsing input number %s (%s)", data, err)
	}
	switch part {
	case 1:
		side := int(math.Ceil(math.Sqrt(float64(num))))
		if side%2 == 0 {
			side++
		}
		prevMax := (side - 2) * (side - 2)
		middle := prevMax + side/2
		x := side / 2
		y := prevMax + (num-prevMax)%(side-1)
		y -= middle
		if y < 0 {
			y = -y
		}
		result := x + y
		log.Printf("Steps required to get data from cell %d: %d", num, result)
	case 2:
		squares := make(map[Point]int)
		x := 0
		y := 0
		side := 0
		squares[Point{0, 0}] = 1
		result := 0
		for {
			x++
			side += 2
			value := squares[Point{x - 1, y + 1}] + squares[Point{x - 1, y}] + squares[Point{x - 1, y - 1}] +
				squares[Point{x, y + 1}] + squares[Point{x, y - 1}] +
				squares[Point{x + 1, y + 1}] + squares[Point{x + 1, y}] + squares[Point{x + 1, y - 1}]
			if value > num {
				result = value
				break
			}
			squares[Point{x, y}] = value
			for i := 1; i < side; i++ {
				y++
				value := squares[Point{x - 1, y + 1}] + squares[Point{x - 1, y}] + squares[Point{x - 1, y - 1}] +
					squares[Point{x, y + 1}] + squares[Point{x, y - 1}] +
					squares[Point{x + 1, y + 1}] + squares[Point{x + 1, y}] + squares[Point{x + 1, y - 1}]
				if value > num {
					result = value
					break
				}
				squares[Point{x, y}] = value
			}
			if result != 0 {
				break
			}
			for i := 0; i < side; i++ {
				x--
				value := squares[Point{x - 1, y + 1}] + squares[Point{x - 1, y}] + squares[Point{x - 1, y - 1}] +
					squares[Point{x, y + 1}] + squares[Point{x, y - 1}] +
					squares[Point{x + 1, y + 1}] + squares[Point{x + 1, y}] + squares[Point{x + 1, y - 1}]
				if value > num {
					result = value
					break
				}
				squares[Point{x, y}] = value
			}
			if result != 0 {
				break
			}
			for i := 0; i < side; i++ {
				y--
				value := squares[Point{x - 1, y + 1}] + squares[Point{x - 1, y}] + squares[Point{x - 1, y - 1}] +
					squares[Point{x, y + 1}] + squares[Point{x, y - 1}] +
					squares[Point{x + 1, y + 1}] + squares[Point{x + 1, y}] + squares[Point{x + 1, y - 1}]
				if value > num {
					result = value
					break
				}
				squares[Point{x, y}] = value
			}
			if result != 0 {
				break
			}
			for i := 0; i < side; i++ {
				x++
				value := squares[Point{x - 1, y + 1}] + squares[Point{x - 1, y}] + squares[Point{x - 1, y - 1}] +
					squares[Point{x, y + 1}] + squares[Point{x, y - 1}] +
					squares[Point{x + 1, y + 1}] + squares[Point{x + 1, y}] + squares[Point{x + 1, y - 1}]
				if value > num {
					result = value
					break
				}
				squares[Point{x, y}] = value
			}
			if result != 0 {
				break
			}
		}
		log.Printf("First square value above %d is %d at (%d,%d)", num, result, x, y)
	}
}
