package main

import (
	"fmt"
	"log"
	"strings"
)

type TuringMachineAction struct {
	writeValue int
	move       int
	nextState  string
}

func NewTuringMachineAction(writeValue int, move int, nextState string) *TuringMachineAction {
	return &TuringMachineAction{
		writeValue: writeValue,
		move:       move,
		nextState:  nextState,
	}
}

type TuringMachineState struct {
	actions map[int]*TuringMachineAction
}

func NewTuringMachineState() *TuringMachineState {
	return &TuringMachineState{
		actions: make(map[int]*TuringMachineAction),
	}
}

func (tms *TuringMachineState) SetAction(value int, action *TuringMachineAction) {
	tms.actions[value] = action
}

type TuringMachine struct {
	tape       map[int]int
	cursor     int
	state      string
	states     map[string]*TuringMachineState
	diagnostic int
}

func NewTuringMachine(initialState string, diagnostic int) *TuringMachine {
	return &TuringMachine{
		tape:       make(map[int]int),
		cursor:     0,
		state:      initialState,
		states:     make(map[string]*TuringMachineState),
		diagnostic: diagnostic,
	}
}

func (tm *TuringMachine) AddState(name string, state *TuringMachineState) {
	tm.states[name] = state
}

func (tm *TuringMachine) Step() {
	action := tm.states[tm.state].actions[tm.tape[tm.cursor]]
	tm.tape[tm.cursor] = action.writeValue
	tm.cursor += action.move
	tm.state = action.nextState
}

func (tm *TuringMachine) Checksum() int {
	result := 0
	for _, v := range tm.tape {
		if v == 1 {
			result++
		}
	}
	return result
}

func (tm *TuringMachine) Run() {
	for i := 0; i < tm.diagnostic; i++ {
		tm.Step()
	}
}

func ParseTuringMachine(data string) *TuringMachine {
	lines := strings.Split(string(data), "\n")
	initialState := lines[0][len(lines[0])-2 : len(lines[0])-1]
	var diagnostic int
	fmt.Sscanf(lines[1], "Perform a diagnostic checksum after %d steps.", &diagnostic)
	result := NewTuringMachine(initialState, diagnostic)

	lines = lines[2:]
	var stateName, nextState, direction string
	var currentValue, writeValue, move int
	for len(lines) > 0 {
		lines = lines[1:]
		fmt.Sscanf(lines[0], "In state %s:", &stateName)
		stateName = strings.Trim(stateName, ":")
		state := NewTuringMachineState()
		fmt.Sscanf(lines[1], "  If the current value is %d:", &currentValue)
		fmt.Sscanf(lines[2], "    - Write the value %d.", &writeValue)
		fmt.Sscanf(lines[3], "    - Move one slot to the %s.", &direction)
		fmt.Sscanf(lines[4], "    - Continue with state %s.", &nextState)
		direction = strings.Trim(direction, ".")
		nextState = strings.Trim(nextState, ".")
		if direction == "left" {
			move = -1
		} else {
			move = 1
		}
		state.SetAction(currentValue, NewTuringMachineAction(writeValue, move, nextState))
		fmt.Sscanf(lines[5], "  If the current value is %d:", &currentValue)
		fmt.Sscanf(lines[6], "    - Write the value %d.", &writeValue)
		fmt.Sscanf(lines[7], "    - Move one slot to the %s.", &direction)
		fmt.Sscanf(lines[8], "    - Continue with state %s.", &nextState)
		direction = strings.Trim(direction, ".")
		nextState = strings.Trim(nextState, ".")
		if direction == "left" {
			move = -1
		} else {
			move = 1
		}
		state.SetAction(currentValue, NewTuringMachineAction(writeValue, move, nextState))
		result.AddState(stateName, state)
		lines = lines[9:]
	}

	return result
}

func Day25(part int, data []byte) {
	turingMachine := ParseTuringMachine(string(data))
	log.Printf("Loaded turing machine with %d states", len(turingMachine.states))
	switch part {
	case 1:
		turingMachine.Run()
		log.Printf("After %d steps checksum of turing machine tape is %d", turingMachine.diagnostic, turingMachine.Checksum())
	case 2:
		log.Printf("Day 25 has no part 2")
	}
}
