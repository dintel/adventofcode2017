package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Day3(part int, data []byte) {
	switch part {
	case 1:
		num, err := strconv.Atoi(strings.TrimSpace(string(data)))
		if err != nil {
			log.Fatalf("Failed parsing input number %s (%s)", data, err)
		}
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
		log.Print("Not implemented yet")
	}
}
