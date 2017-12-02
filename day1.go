package main

import (
	"log"
)

func Day1(part int, data []byte) {
	switch part {
	case 1:
		data = append(data, data[0])
		result := 0
		for i := 0; i < len(data)-1; i++ {
			if data[i] == data[i+1] {
				result += int(data[i] - '0')
			}
		}
		log.Printf("Solution to captcha is %d", result)
	case 2:
	}
}
