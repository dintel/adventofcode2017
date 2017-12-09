package main

import (
	"bytes"
	"log"
	"regexp"
	"strings"
)

func Day9(part int, data []byte) {
	//input := string(data)
	switch part {
	case 1:
		excl := false
		filteredBuffer := bytes.NewBuffer(nil)
		for _, ch := range data {
			if excl {
				excl = false
				continue
			}
			if ch == '!' {
				excl = true
				continue
			}
			filteredBuffer.WriteByte(ch)
		}
		filteredInput := filteredBuffer.String()
		garbage := regexp.MustCompile("<.*?>")
		filteredInput = garbage.ReplaceAllString(filteredInput, "")
		filteredInput = strings.Replace(filteredInput, ",", "", -1)
		currentNesting := 0
		totalScore := 0
		for _, ch := range filteredInput {
			switch ch {
			case '{':
				currentNesting++
				totalScore += currentNesting
			case '}':
				currentNesting--
			default:
				log.Fatalf("unknown character encountered (%c)", ch)
			}
		}
		log.Printf("Total score is %d", totalScore)
	case 2:
		excl := false
		filteredBuffer := bytes.NewBuffer(nil)
		for _, ch := range data {
			if excl {
				excl = false
				continue
			}
			if ch == '!' {
				excl = true
				continue
			}
			filteredBuffer.WriteByte(ch)
		}
		filteredInput := filteredBuffer.String()
		garbageRe := regexp.MustCompile("<.*?>")
		garbage := garbageRe.FindAllString(filteredInput, -1)
		totalGarbage := 0
		for _, part := range garbage {
			totalGarbage += len(part) - 2
		}
		log.Printf("Found %d garbage", totalGarbage)
	}
}
