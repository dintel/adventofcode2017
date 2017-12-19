package main

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const DUET_BUFFER_SIZE = 1000

type DeadlockDetector struct {
	mutex sync.Mutex
	rcv   map[int]bool
	ch    chan bool
}

func NewDeadlockDetector() *DeadlockDetector {
	return &DeadlockDetector{
		rcv: make(map[int]bool),
		ch:  make(chan bool),
	}
}

func (d *DeadlockDetector) Ch() chan bool {
	return d.ch
}

func (d *DeadlockDetector) Set(pid int, rcv bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.rcv[pid] = rcv
}

func (d *DeadlockDetector) Deadlock() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	result := true
	for _, v := range d.rcv {
		result = result && v
	}
	return result
}

func (d *DeadlockDetector) Run() {
	timer := time.NewTimer(time.Millisecond * 50)
	counter := 0
	for _ = range timer.C {
		if d.Deadlock() {
			counter++
		} else {
			counter = 0
		}
		if counter == 2 {
			d.ch <- true
			break
		}
		timer.Reset(time.Millisecond * 50)
	}
}

type DuetInstruction struct {
	operation string
	args      []string
}

type DuetProcessor struct {
	pid          int
	registers    map[string]int
	program      []DuetInstruction
	current      int
	sendCh       chan int
	recvCh       chan int
	lastRecieved int
	sendCounter  int
	dd           *DeadlockDetector
}

func NewDuetProcessor(program []DuetInstruction, pid int, dd *DeadlockDetector) *DuetProcessor {
	result := &DuetProcessor{
		pid:         pid,
		registers:   make(map[string]int),
		program:     program,
		current:     0,
		sendCh:      nil,
		recvCh:      make(chan int, DUET_BUFFER_SIZE),
		sendCounter: 0,
		dd:          dd,
	}
	result.Set("p", pid)
	return result
}

func (duet *DuetProcessor) Run(stopAfter string) {
	if stopAfter != "" {
		for duet.current < len(duet.program) && duet.Exec() != stopAfter {
		}
	} else {
		for duet.current < len(duet.program) {
			duet.Exec()
		}
	}
}

func (duet *DuetProcessor) Exec() string {
	i := duet.program[duet.current]
	duet.current++
	switch i.operation {
	case "snd":
		duet.sendCounter++
		duet.sendCh <- duet.Value(i.args[0])
	case "set":
		r := i.args[0]
		v := duet.Value(i.args[1])
		duet.Set(r, v)
	case "add":
		r := i.args[0]
		v := duet.Value(i.args[1])
		duet.Set(r, duet.Get(r)+v)
	case "mul":
		r := i.args[0]
		v := duet.Value(i.args[1])
		duet.Set(r, duet.Get(r)*v)
	case "mod":
		r := i.args[0]
		v := duet.Value(i.args[1])
		duet.Set(r, duet.Get(r)%v)
	case "rcv":
		duet.dd.Set(duet.pid, true)
		v := <-duet.recvCh
		duet.dd.Set(duet.pid, false)
		duet.Set(i.args[0], v)
		duet.lastRecieved = v
	case "jgz":
		cond := duet.Value(i.args[0])
		offset := duet.Value(i.args[1])
		if cond > 0 {
			duet.current--
			duet.current += offset
		}
	default:
		log.Fatalf("Unknown operation '%s'", i.operation)
	}
	return i.operation
}

func (duet *DuetProcessor) Value(arg string) int {
	result, err := strconv.Atoi(arg)
	if err != nil {
		return duet.registers[arg]
	}
	return result
}

func (duet *DuetProcessor) Set(register string, value int) {
	duet.registers[register] = value
}

func (duet *DuetProcessor) SetProgram(program []DuetInstruction) {
	duet.program = program
}

func (duet *DuetProcessor) Get(register string) int {
	return duet.registers[register]
}

func (duet *DuetProcessor) SetSendChannel(ch chan int) {
	duet.sendCh = ch
}

func (duet *DuetProcessor) RecvChannel() chan int {
	return duet.recvCh
}

func (duet *DuetProcessor) LastRecieved() int {
	for len(duet.recvCh) > 0 {
		duet.lastRecieved = <-duet.recvCh
	}
	return duet.lastRecieved
}

func (duet *DuetProcessor) SendCounter() int {
	return duet.sendCounter
}

func Day18(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	program := make([]DuetInstruction, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		program[i].operation = fields[0]
		program[i].args = fields[1:]
	}
	log.Printf("Loaded program that consists of %d instructions", len(program))
	deadlockDetector := NewDeadlockDetector()
	proc0 := NewDuetProcessor(program, 0, deadlockDetector)
	proc1 := NewDuetProcessor(program, 1, deadlockDetector)
	proc0.SetSendChannel(proc1.RecvChannel())
	proc1.SetSendChannel(proc0.RecvChannel())
	switch part {
	case 1:
		proc0.SetSendChannel(proc0.RecvChannel())
		proc0.Run("rcv")
		for proc0.LastRecieved() == 0 {
			proc0.Run("rcv")
		}
		log.Printf("Duet program value after first rcv is %d", proc0.LastRecieved())
	case 2:
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			<-deadlockDetector.Ch()
			log.Print("Deadlock detected")
			wg.Done()
			wg.Done()
		}()
		go func() {
			deadlockDetector.Run()
		}()
		go func() {
			proc0.Run("")
			wg.Done()
		}()
		go func() {
			proc1.Run("")
			wg.Done()
		}()
		wg.Wait()
		log.Printf("Program 1 send counter is %d", proc1.SendCounter())
	}
}
