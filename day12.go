package main

import (
	"log"
	"strconv"
	"strings"
)

func scanGroup(connections map[int][]int, root int) map[int]bool {
	visited := make(map[int]bool)
	queue := make([]int, 1)
	queue[0] = root
	var current int
	for len(queue) != 0 {
		current, queue = queue[0], queue[1:]
		visited[current] = true
		for _, neighbor := range connections[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return visited
}

func Day12(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	connections := make(map[int][]int)
	edgeCounter := 0
	for i, line := range lines {
		line = strings.Replace(line, ",", " ", -1)
		line = strings.Replace(line, "<->", " ", -1)
		fields := strings.Fields(line)
		if len(fields) == 0 {
			log.Fatalf("Failed parsing line %d", i)
		}
		src, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("Failed parsing number %s at line %d", fields[0], i)
		}
		for j := 1; j < len(fields); j++ {
			dst, err := strconv.Atoi(fields[j])
			if err != nil {
				log.Fatalf("Failed parsing number %s at line %d", fields[j], i)
			}
			connections[src] = append(connections[src], dst)
			edgeCounter++
		}
	}
	log.Printf("Loaded %d vertices, %d edges", len(connections), edgeCounter)
	switch part {
	case 1:
		visited := scanGroup(connections, 0)
		log.Printf("Tree with root 0 has %d vertices", len(visited))
	case 2:
		groups := 0
		visited := make(map[int]bool)
		for v := range connections {
			visited[v] = false
		}
		for v := range visited {
			if !visited[v] {
				group := scanGroup(connections, v)
				groups++
				for groupVertex := range group {
					visited[groupVertex] = true
				}
			}
		}
		log.Printf("There are %d groups", groups)
	}
}
