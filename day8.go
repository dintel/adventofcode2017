package main

import (
	"log"
	"strconv"
	"strings"
)

type Processor struct {
	Registers map[string]int
}

func NewProcessor() *Processor {
	return &Processor{Registers: make(map[string]int)}
}

func (p *Processor) Inc(register string, amount int) {
	p.Registers[register] += amount
}

func (p *Processor) Dec(register string, amount int) {
	p.Registers[register] -= amount
}

func (p *Processor) Check(cond Condition) bool {
	left := p.Registers[cond.Left]
	right := cond.Right
	switch cond.Operator {
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">":
		return left > right
	case ">=":
		return left >= right
	case "<":
		return left < right
	case "<=":
		return left <= right
	}
	log.Fatalf("unknown operator '%s' in condition", cond.Operator)
	return false
}

func (p *Processor) Execute(operation string, register string, operand int) {
	switch operation {
	case "inc":
		p.Inc(register, operand)
		return
	case "dec":
		p.Dec(register, operand)
		return
	}
	log.Fatalf("unknown operation '%s' in instruction", operation)
}

func (p *Processor) Max() (string, int) {
	maxR := ""
	maxValue := 0
	for r, v := range p.Registers {
		if maxValue < v {
			maxValue = v
			maxR = r
		}
	}
	return maxR, maxValue
}

type Condition struct {
	Left     string
	Operator string
	Right    int
}

type Instruction struct {
	Register  string
	Operation string
	Operand   int
	Condition Condition
}

func Day8(part int, data []byte) {
	var err error
	lines := strings.Split(string(data), "\n")
	processor := NewProcessor()
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		instructions[i].Register = fields[0]
		instructions[i].Operation = fields[1]
		instructions[i].Operand, err = strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalf("Failed parsing number %s at line %d", fields[2], i)
		}
		instructions[i].Condition.Left = fields[4]
		instructions[i].Condition.Operator = fields[5]
		instructions[i].Condition.Right, err = strconv.Atoi(fields[6])
		if err != nil {
			log.Fatalf("Failed parsing number %s at line %d", fields[2], i)
		}
	}
	log.Printf("Loaded %d instructions", len(instructions))
	switch part {
	case 1:
		for _, instruction := range instructions {
			if processor.Check(instruction.Condition) {
				processor.Execute(instruction.Operation, instruction.Register, instruction.Operand)
			}
		}

		maxR, maxValue := processor.Max()
		log.Printf("Largets value is in register %s=%d", maxR, maxValue)
	case 2:
		maxR, maxValue := "", 0
		for _, instruction := range instructions {
			if processor.Check(instruction.Condition) {
				processor.Execute(instruction.Operation, instruction.Register, instruction.Operand)
				newMaxR, newMaxValue := processor.Max()
				if newMaxValue > maxValue {
					maxR, maxValue = newMaxR, newMaxValue
				}
			}
		}

		log.Printf("Largets value during execution was in register %s=%d", maxR, maxValue)
	}
}
