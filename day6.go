package main

import (
	"log"
	"strconv"
	"strings"
)

func cmpBanks(bank1, bank2 []int) bool {
	for i := range bank1 {
		if bank1[i] != bank2[i] {
			return false
		}
	}
	return true
}

func containsBank(history [][]int, bank []int) int {
	for i := range history {
		if cmpBanks(history[i], bank) {
			return i
		}
	}
	return -1
}

func shuffle(banks []int) {
	idx := 0
	for i := range banks {
		if banks[i] > banks[idx] {
			idx = i
		}
	}

	surplus := banks[idx]/len(banks) + 1
	firstMore := banks[idx] % len(banks)
	banks[idx] = 0
	for i := 0; i < len(banks); i++ {
		idx++
		if idx >= len(banks) {
			idx -= len(banks)
		}
		if i >= firstMore {
			i += len(banks)
			surplus--
		}
		banks[idx] += surplus
	}
}

func historyAppend(history [][]int, banks []int) [][]int {
	banksCopy := make([]int, len(banks))
	copy(banksCopy, banks)
	return append(history, banksCopy)
}

func Day6(part int, data []byte) {
	var err error
	fields := strings.Fields(string(data))
	banks := make([]int, len(fields))
	for i, field := range fields {
		banks[i], err = strconv.Atoi(field)
		if err != nil {
			log.Fatalf("Failed parsing number %s at field %d", field, i)
		}
	}
	log.Printf("Loaded %d banks: %v", len(banks), banks)
	history := make([][]int, 1)
	history[0] = make([]int, len(banks))
	copy(history[0], banks)
	switch part {
	case 1:
		for {
			shuffle(banks)
			if containsBank(history, banks) != -1 {
				break
			}
			history = historyAppend(history, banks)
		}
		log.Printf("detected loop after %d steps", len(history))
	case 2:
		loopSize := 0
		for {
			shuffle(banks)
			duplicate := containsBank(history, banks)
			if duplicate != -1 {
				loopSize = len(history) - duplicate
				break
			}
			history = historyAppend(history, banks)
		}
		log.Printf("detected loop size %d", loopSize)
	}
}
