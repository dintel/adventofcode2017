package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

const COPROCESSOR_STAT_OP = "mul"
const COPROCESSOR_REGISTER = "h"
const COPROCESSOR_INIT = 7

type CoprocessorInstruction struct {
	operation string
	args      []string
}

type Coprocessor struct {
	registers map[string]int
	program   []CoprocessorInstruction
	stat      map[string]int
	current   int
}

func NewCoprocessor(program []CoprocessorInstruction) *Coprocessor {
	result := &Coprocessor{
		registers: make(map[string]int),
		program:   program,
		stat:      make(map[string]int),
		current:   0,
	}
	return result
}

func (cop *Coprocessor) Run() {
	for cop.current >= 0 && cop.current < len(cop.program) {
		cop.Exec()
	}
}

func (cop *Coprocessor) Exec() string {
	i := cop.program[cop.current]
	cop.current++
	cop.stat[i.operation]++
	switch i.operation {
	case "set":
		r := i.args[0]
		v := cop.Value(i.args[1])
		cop.Set(r, v)
	case "sub":
		r := i.args[0]
		v := cop.Value(i.args[1])
		cop.Set(r, cop.Get(r)-v)
	case "mul":
		r := i.args[0]
		v := cop.Value(i.args[1])
		cop.Set(r, cop.Get(r)*v)
	case "jnz":
		cond := cop.Value(i.args[0])
		offset := cop.Value(i.args[1])
		if cond != 0 {
			cop.current--
			cop.current += offset
		}
	default:
		log.Fatalf("Unknown operation '%s'", i.operation)
	}
	return i.operation
}

func (cop *Coprocessor) Value(arg string) int {
	result, err := strconv.Atoi(arg)
	if err != nil {
		return cop.registers[arg]
	}
	return result
}

func (cop *Coprocessor) Set(register string, value int) {
	cop.registers[register] = value
}

func (cop *Coprocessor) Get(register string) int {
	return cop.registers[register]
}

func (cop *Coprocessor) Stat(op string) int {
	return cop.stat[op]
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func Day23(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	program := make([]CoprocessorInstruction, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		program[i].operation = fields[0]
		program[i].args = fields[1:]
	}
	log.Printf("Loaded program that consists of %d instructions", len(program))
	coprocessor := NewCoprocessor(program)
	switch part {
	case 1:
		coprocessor.Run()
		log.Printf("after finishing program execution instruction %s was executed %d times", COPROCESSOR_STAT_OP, coprocessor.Stat(COPROCESSOR_STAT_OP))
	case 2:
		// It's a cheating here. Instead of code optimization, we just calculate
		// what program does.
		coprocessor.Set("a", 1)
		for i := 0; i < COPROCESSOR_INIT; i++ {
			coprocessor.Exec()
		}
		current := coprocessor.Get("b")
		h := 0
		for i := 0; i <= 1000; i++ {
			if !IsPrime(current) {
				h++
			}
			current += 17
		}
		log.Printf("after finishing program execution register %s has value %d", COPROCESSOR_REGISTER, h)
	}
}
