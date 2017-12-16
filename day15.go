package main

import (
	"fmt"
	"log"
	"strings"
)

const PAIRS_TO_CHECK = 40000000
const PAIRS_TO_CHECK2 = 5000000
const GENERATOR_DIVIDER = 2147483647

type Generator struct {
	current int
	factor  int
	divisor int
}

func NewGenerator(firstNum int, factor int, divisor int) *Generator {
	return &Generator{
		current: firstNum,
		factor:  factor,
		divisor: divisor,
	}
}

func (g *Generator) Next() int {
	g.current = (g.current * g.factor) % GENERATOR_DIVIDER
	return g.current
}

func (g *Generator) Next2() int {
	next := g.Next()
	for next%g.divisor != 0 {
		next = g.Next()
	}
	return next
}

func Day15(part int, data []byte) {
	var err error
	lines := strings.Split(string(data), "\n")
	factors := map[string]int{
		"A": 16807,
		"B": 48271,
	}
	divisors := map[string]int{
		"A": 4,
		"B": 8,
	}
	generators := make(map[string]*Generator, len(lines))
	for i, line := range lines {
		var name string
		var first int
		_, err = fmt.Sscanf(line, "Generator %s starts with %d", &name, &first)
		if err != nil {
			log.Fatalf("Failed parsing line %d", i)
		}
		generators[name] = NewGenerator(first, factors[name], divisors[name])
	}
	if generators["A"] == nil {
		log.Fatal("Generator A not found")
	}
	if generators["B"] == nil {
		log.Fatal("Generator B not found")
	}
	switch part {
	case 1:
		matches := 0
		for i := 0; i < PAIRS_TO_CHECK; i++ {
			a := generators["A"].Next()
			b := generators["B"].Next()
			if a&0xffff == b&0xffff {
				matches++
			}
		}
		log.Printf("Found %d matches in %d pairs", matches, PAIRS_TO_CHECK)
	case 2:
		matches := 0
		for i := 0; i < PAIRS_TO_CHECK2; i++ {
			a := generators["A"].Next2()
			b := generators["B"].Next2()
			if a&0xffff == b&0xffff {
				matches++
			}
		}
		log.Printf("Found %d matches in %d pairs", matches, PAIRS_TO_CHECK2)
	}
}
