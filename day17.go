package main

import (
	"log"
	"strconv"
)

const SPINLOCK_STEPS = 2017
const SPINLOCK_STEPS2 = 50000000

type CircularNode struct {
	next  *CircularNode
	value int
}

type Spinlock struct {
	first     *CircularNode
	current   *CircularNode
	step      int
	nextValue int
}

func NewSpinlock(step int) *Spinlock {
	first := &CircularNode{
		value: 0,
	}
	first.next = first
	return &Spinlock{
		first:     first,
		current:   first,
		step:      step,
		nextValue: 1,
	}
}

func (s *Spinlock) Step() {
	for i := 0; i < s.step; i++ {
		s.current = s.current.next
	}
	node := &CircularNode{
		next:  s.current.next,
		value: s.nextValue,
	}
	s.current.next = node
	s.current = s.current.next
	s.nextValue++
}

func (s *Spinlock) AfterNext() int {
	return s.current.next.value
}

func (s *Spinlock) AfterFirst() int {
	return s.first.next.value
}

func Day17(part int, data []byte) {
	step, err := strconv.Atoi(string(data))
	if err != nil {
		log.Fatalf("Failed parsing spinlock step %s", string(data))
	}
	spinlock := NewSpinlock(step)
	log.Printf("Spinlock step is %d", step)
	switch part {
	case 1:
		for i := 0; i < SPINLOCK_STEPS; i++ {
			spinlock.Step()
		}
		log.Printf("Value after last insertion is %d", spinlock.AfterNext())
	case 2:
		log.Print("Allocating one by one")
		nodes := make([]*CircularNode, SPINLOCK_STEPS2)
		for i := 0; i < SPINLOCK_STEPS2; i++ {
			nodes[i] = new(CircularNode)
		}
		log.Print("Done")
		log.Print("Allocating all at once")
		nodes2 := make([]CircularNode, SPINLOCK_STEPS2)
		nodes2[0].next = nil
		log.Print("Done")
		percent := 0
		for i := 0; i < SPINLOCK_STEPS2; i++ {
			if i%500000 == 0 && i != 0 {
				percent++
				log.Printf("%d%% completed", percent)
			}
			spinlock.Step()
		}
		log.Printf("Value at position 1 is %d", spinlock.AfterFirst())
	}
}
