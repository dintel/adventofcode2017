package main

import (
	"log"
	"strconv"
	"strings"
)

func tripSeverity(scanners map[int]int) int {
	severity := 0
	for layer, depth := range scanners {
		mod := (depth - 1) * 2
		if layer%mod == 0 {
			severity += layer * depth
		}
	}
	return severity
}

func safeTrip(scanners map[int]int, delay int) bool {
	for layer, depth := range scanners {
		mod := (depth - 1) * 2
		if (layer+delay)%mod == 0 {
			return false
		}
	}
	return true
}

func Day13(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	scanners := make(map[int]int)
	maxLayer := 0
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			log.Fatalf("Failed parsing line %d - %s", i, line)
		}
		layer, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed parsing layer number %s at line %d", parts[0], i)
		}
		depth, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed parsing scanner depth %s at line %d", parts[1], i)
		}
		scanners[layer] = depth
		if layer > maxLayer {
			maxLayer = layer
		}
	}
	log.Printf("Loaded %d scanners, max layer number is %d", len(scanners), maxLayer)
	switch part {
	case 1:
		log.Printf("Trip severity is %d", tripSeverity(scanners))
	case 2:
		delay := 0
		for !safeTrip(scanners, delay) {
			delay++
		}
		log.Printf("Earliest safe trip is with delay of %d picoseconds", delay)
	}
}
