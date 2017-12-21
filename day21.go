package main

import (
	"log"
	"strings"
)

const PICTURE_TRANSFORMS = 5
const PICTURE_TRANSFORMS2 = 18

type Square [][]int

func NewSquare(size int) *Square {
	result := Square(make([][]int, size))
	for i := range result {
		result[i] = make([]int, size)
	}
	return &result
}

func ParseSquare(s string) *Square {
	lines := strings.Split(s, "/")
	size := len(lines)
	square := Square(make([][]int, size))
	for i := range square {
		square[i] = make([]int, size)
		for j := range square[i] {
			if lines[i][j] == '.' {
				square[i][j] = 0
			} else {
				square[i][j] = 1
			}
		}
	}
	return &square
}

func (s *Square) Rotate() {
	for i := 0; i < len(*s); i++ {
		for j := i; j < len(*s); j++ {
			(*s)[i][j], (*s)[j][i] = (*s)[j][i], (*s)[i][j]
		}
	}
	s.Flip(false)
}

func (s *Square) Flip(horizontal bool) {
	half := len(*s) / 2
	if horizontal {
		for i := 0; i < half; i++ {
			(*s)[i], (*s)[len(*s)-1-i] = (*s)[len(*s)-1-i], (*s)[i]
		}
	} else {
		for i := 0; i < len(*s); i++ {
			for j := 0; j < half; j++ {
				(*s)[i][j], (*s)[i][len(*s)-1-j] = (*s)[i][len(*s)-1-j], (*s)[i][j]
			}
		}
	}
}

func (s *Square) Int() int {
	p := 1
	result := 0
	for i := range *s {
		for j := range (*s)[i] {
			result += (*s)[i][j] * p
			p *= 2
		}
	}
	if len(*s) == 2 {
		result += 1024
	}
	return result
}

func (s *Square) String() string {
	result := ""
	for _, row := range *s {
		for _, v := range row {
			if v == 0 {
				result += "."
			} else {
				result += "#"
			}
		}
		result += "\n"
	}
	return strings.Trim(result, "\n")
}

func (s *Square) sub(x int, y int, size int) *Square {
	part := Square(make([][]int, size))
	for i := range part {
		part[i] = (*s)[y+i][x : x+size]
	}
	return &part
}

func (s *Square) copy(x int, y int, src *Square) {
	for i := range *src {
		for j := range (*src)[i] {
			(*s)[y+i][x+j] = (*src)[i][j]
		}
	}
}

func (s *Square) Transform(rules map[int]*Square) *Square {
	var srcSize int
	if len(*s)%2 == 0 {
		srcSize = 2
	} else if len(*s)%3 == 0 {
		srcSize = 3
	} else {
		log.Fatalf("Square size %d is not divisible by 2 and 3", len(*s))
	}
	tgtSize := srcSize + 1
	parts := len(*s) / srcSize
	picture := NewSquare(parts * tgtSize)
	for i := 0; i < parts; i++ {
		for j := 0; j < parts; j++ {
			part := s.sub(i*srcSize, j*srcSize, srcSize)
			id := part.Int()
			if rules[id] == nil {
				log.Fatalf("No rule found for square ID %d\n%s", id, part)
			}
			picture.copy(i*tgtSize, j*tgtSize, rules[id])
		}
	}
	return picture
}

func (s *Square) Count(v int) int {
	result := 0
	for i := range *s {
		for j := range (*s)[i] {
			if (*s)[i][j] == v {
				result++
			}
		}
	}
	return result
}

func Day21(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	rules := make(map[int]*Square)
	for i, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			log.Fatalf("Failed parsing rule at line %d", i)
		}
		from := ParseSquare(parts[0])
		to := ParseSquare(parts[1])

		rules[from.Int()] = to
		from.Flip(true)
		rules[from.Int()] = to
		from.Flip(true)
		from.Flip(false)
		rules[from.Int()] = to
		from.Flip(false)

		from.Rotate()
		rules[from.Int()] = to
		from.Flip(true)
		rules[from.Int()] = to
		from.Flip(true)
		from.Flip(false)
		rules[from.Int()] = to
		from.Flip(false)

		from.Rotate()
		rules[from.Int()] = to
		from.Flip(true)
		rules[from.Int()] = to
		from.Flip(true)
		from.Flip(false)
		rules[from.Int()] = to
		from.Flip(false)

		from.Rotate()
		rules[from.Int()] = to
		from.Flip(true)
		rules[from.Int()] = to
		from.Flip(true)
		from.Flip(false)
		rules[from.Int()] = to
		from.Flip(false)

	}
	log.Printf("Parsed %d rules, created %d rules", len(lines), len(rules))
	picture := &Square{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}}
	switch part {
	case 1:
		for i := 0; i < PICTURE_TRANSFORMS; i++ {
			picture = picture.Transform(rules)
		}
		log.Printf("After %d iterations there are %d pixels on", PICTURE_TRANSFORMS, picture.Count(1))
	case 2:
		for i := 0; i < PICTURE_TRANSFORMS2; i++ {
			picture = picture.Transform(rules)
		}
		log.Printf("After %d iterations there are %d pixels on", PICTURE_TRANSFORMS2, picture.Count(1))
	}
}
