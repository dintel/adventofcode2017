package main

import (
	"log"
	"strconv"
	"strings"
)

func Day2(part int, data []byte) {
	switch part {
	case 1:
		lines := strings.Split(string(data), "\n")
		result := 0
		for i, line := range lines {
			row := strings.Fields(line)
			rowMax, rowMin := 0, 0
			for j, term := range row {
				num, err := strconv.Atoi(term)
				if err != nil {
					log.Fatalf("Failed parsing row %d, cell %d (%s)", i, j, err)
				}
				if rowMax == 0 || num > rowMax {
					rowMax = num
				}
				if rowMin == 0 || num < rowMin {
					rowMin = num
				}
			}
			result += rowMax - rowMin
		}
		log.Printf("Checksum is %d", result)
	case 2:
		log.Printf("Not implemented yet")
	}
}
