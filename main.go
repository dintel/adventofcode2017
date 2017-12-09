package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <day-number> <part-number> parameters\n", os.Args[0])
	}

	// Parse day number and validate it
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Could not parse day number: %s", err)
	}
	if day < 1 || day > 25 {
		log.Fatalln("Day number must be between 1 and 25")
	}

	// Parse part number and validate it
	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Could not parse part number: %s", err)
	}
	if part < 1 || part > 2 {
		log.Fatalln("Part number must be erither 1 or 2")
	}

	// Load input data
	if len(os.Args) != 4 {
		log.Fatalf("Expected input file parameter")
	}
	filename := os.Args[3]
	log.Printf("Loading file %s", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read input file - %s", err)
	}

	switch day {
	case 1:
		Day1(part, data)
	case 2:
		Day2(part, data)
	case 3:
		Day3(part, data)
	case 4:
		Day4(part, data)
	case 5:
		Day5(part, data)
	case 6:
		Day6(part, data)
	case 7:
		Day7(part, data)
	case 8:
		Day8(part, data)
	case 9:
		Day9(part, data)
	default:
		log.Printf("Day %d not implemented yet", day)
	}
	end := time.Now()
	duration := end.Sub(start)
	log.Printf("Runtime: %s", duration)
}
