package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Port [2]int

func ParsePort(s string) (*Port, error) {
	port := new(Port)
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return nil, errors.New(fmt.Sprintf("failed parsing connector '%s'", s))
	}
	var err error
	(*port)[0], err = strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	(*port)[1], err = strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return port, nil
}

func (p *Port) String() string {
	return fmt.Sprintf("%d/%d", (*p)[0], (*p)[1])
}

func (p *Port) Flip() {
	p[0], p[1] = p[1], p[0]
}

type Bridge []*Port

func (b *Bridge) Value() int {
	result := 0
	for _, port := range *b {
		result += (*port)[0] + (*port)[1]
	}
	return result
}

func PortsWithout(ports []*Port, excl int) []*Port {
	result := make([]*Port, len(ports)-1)
	j := 0
	for i := range ports {
		if i == excl {
			continue
		}
		result[j] = ports[i]
		j++
	}
	return result
}

func FindBridges(ports []*Port, part Bridge) []Bridge {
	result := make([]Bridge, 1)
	result[0] = part
	start := 0
	if len(part) != 0 {
		start = part[len(part)-1][1]
	}
	for i, port := range ports {
		if port[1] == start {
			port.Flip()
		}
		if port[0] == start {
			newPorts := PortsWithout(ports, i)
			newPart := Bridge(make([]*Port, len(part)+1))
			copy(newPart, part)
			newPart[len(newPart)-1] = port
			result = append(result, FindBridges(newPorts, newPart)...)
		}
	}
	return result
}

func Day24(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	ports := make([]*Port, len(lines))
	for i, line := range lines {
		port, err := ParsePort(line)
		if err != nil {
			log.Fatalf("failed parsing line %d - %s", i, err)
		}
		ports[i] = port
	}
	log.Printf("Loaded %d connectors", len(ports))
	switch part {
	case 1:
		best := Bridge(nil)
		for _, bridge := range FindBridges(ports, nil) {
			if best.Value() < bridge.Value() {
				best = bridge
			}
		}
		log.Printf("best bridge found has value of %d", best.Value())
	case 2:
		best := Bridge(nil)
		for _, bridge := range FindBridges(ports, nil) {
			if best.Value() < bridge.Value() && len(best) <= len(bridge) {
				best = bridge
			}
		}
		log.Printf("longest best bridge found has value of %d and has %d connectors", best.Value(), len(best))
	}
}
