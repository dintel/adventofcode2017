package main

import (
	"log"
	"sort"
	"strings"
)

func hasDuplicates(words []string) bool {
	encountered := make(map[string]bool)
	for _, word := range words {
		if encountered[word] {
			return true
		}
		encountered[word] = true
	}
	return false
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func SortStrings(strings []string) []string {
	for i := range strings {
		strings[i] = SortString(strings[i])
	}
	return strings
}

func Day4(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	log.Printf("Loaded %d passphrases", len(lines))
	switch part {
	case 1:
		valid := 0
		for _, line := range lines {
			if !hasDuplicates(strings.Fields(line)) {
				valid++
			}
		}
		log.Printf("Found %d valid passphrases", valid)
	case 2:
		valid := 0
		for _, line := range lines {
			words := SortStrings(strings.Fields(line))
			if !hasDuplicates(words) {
				valid++
			}
		}
		log.Printf("Found %d valid passphrases", valid)
	}
}
