package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const KnotHashSize = 256
const KnotHashBlock = 16
const KnotHashBlocks = 16

type KnotHasher struct {
	list    []int
	current int
	skip    int
}

func NewKnotHasher() *KnotHasher {
	result := &KnotHasher{list: make([]int, KnotHashSize)}
	for i := 0; i < KnotHashSize; i++ {
		result.list[i] = i
	}
	return result
}

func (hasher *KnotHasher) String() string {
	result := ""
	for i := 0; i < KnotHashBlocks; i++ {
		blockValue := 0
		for j := 0; j < KnotHashBlock; j++ {
			blockValue ^= hasher.Get(i*KnotHashBlock + j)
		}
		result += fmt.Sprintf("%02x", blockValue)
	}
	return result
}

func (hasher *KnotHasher) Get(i int) int {
	return hasher.list[i]
}

func (hasher *KnotHasher) Hash(input []int) {
	for _, n := range input {
		hasher.reverse(hasher.current, n)
		hasher.current = (hasher.current + n + hasher.skip) % len(hasher.list)
		hasher.skip = (hasher.skip + 1) % len(hasher.list)
	}
}

func (hasher *KnotHasher) FullHash(data []byte) {
	input := make([]int, len(data))
	for i, ch := range data {
		input[i] = int(ch)
	}
	input = append(input, 17, 31, 73, 47, 23)
	for i := 0; i < 64; i++ {
		hasher.Hash(input)
	}
}

func (hasher *KnotHasher) reverse(current int, length int) {
	tmp := make([]int, len(hasher.list))
	copy(tmp, hasher.list)
	end := current + length - 1
	for i := 0; i < length; i++ {
		src := (end - i) % len(hasher.list)
		hasher.list[current] = tmp[src]
		current++
		current %= len(hasher.list)
	}
}

func Day10(part int, data []byte) {
	var err error
	switch part {
	case 1:
		parts := strings.Split(string(data), ",")
		input := make([]int, len(parts))
		for i, part := range parts {
			input[i], err = strconv.Atoi(part)
			if err != nil {
				log.Fatalf("Failed parsing number %s at position %d", part, i)
			}
		}
		log.Printf("Loaded %d numbers", len(input))
		hasher := NewKnotHasher()
		hasher.Hash(input)
		result := hasher.Get(0) * hasher.Get(1)
		log.Printf("multiplication of first 2 numbers after one round of Knot Hash is %d", result)
	case 2:
		hasher := NewKnotHasher()
		input := make([]int, len(data))
		for i, ch := range data {
			input[i] = int(ch)
		}
		input = append(input, 17, 31, 73, 47, 23)
		for i := 0; i < 64; i++ {
			hasher.Hash(input)
		}
		log.Printf("Knot Hash is %s", hasher)
	}
}
