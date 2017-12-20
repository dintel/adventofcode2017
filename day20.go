package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Vector [3]int

func (v *Vector) add(other Vector) {
	v[0] += other[0]
	v[1] += other[1]
	v[2] += other[2]
}

type Particle struct {
	p Vector
	v Vector
	a Vector
}

func (p *Particle) A() float64 {
	ax := float64(p.a[0])
	ay := float64(p.a[1])
	az := float64(p.a[2])
	return math.Sqrt(ax*ax + ay*ay + az*az)
}

func (p *Particle) Distance(other *Particle) float64 {
	dx := float64(p.p[0] - other.p[0])
	dy := float64(p.p[1] - other.p[1])
	dz := float64(p.p[2] - other.p[2])
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p *Particle) Tick() {
	p.v.add(p.a)
	p.p.add(p.v)
}

func (p *Particle) V() float64 {
	vx := float64(p.v[0])
	vy := float64(p.v[1])
	vz := float64(p.v[2])
	return math.Sqrt(vx*vx + vy*vy + vz*vz)
}

func (p *Particle) P() Vector {
	return p.p
}

type Pair [2]int

type ParticleSystem struct {
	particles map[int]*Particle
	distances map[Pair]float64
	speeds    map[int]float64
}

func NewParticleSystem(particles []*Particle) *ParticleSystem {
	particlesMap := make(map[int]*Particle)
	speeds := make(map[int]float64)
	for i, p := range particles {
		particlesMap[i] = p
		speeds[i] = p.V()
	}
	distances := make(map[Pair]float64)
	for i := 0; i < len(particles); i++ {
		for j := i + 1; j < len(particles); j++ {
			distances[Pair{i, j}] = particles[i].Distance(particles[j])
		}
	}
	return &ParticleSystem{
		particles: particlesMap,
		distances: distances,
		speeds:    speeds,
	}
}

func (ps *ParticleSystem) HasPossiblyColliding() bool {
	return len(ps.distances) != 0
}

func (ps *ParticleSystem) Particles() int {
	return len(ps.particles)
}

func (ps *ParticleSystem) Tick() {
	// Tick all particles and record collisions
	collisions := make(map[Vector][]int)
	for i, p := range ps.particles {
		p.Tick()
		collisions[p.P()] = append(collisions[p.P()], i)
	}

	// Delete collided particles
	for _, collision := range collisions {
		if len(collision) > 1 {
			for _, i := range collision {
				delete(ps.particles, i)
			}
		}
	}

	// Calculate speeds for all particles and write whether they accelerate.
	accelerates := make(map[int]bool)
	for i, p := range ps.particles {
		newV := p.V()
		accelerates[i] = ps.speeds[i] < newV
		ps.speeds[i] = newV
	}

	// Update pair distances. Remove safe pairs, such that both
	// particles accelerate and their distance increased.
	safePairs := make([]Pair, 0)
	for pair, prevDist := range ps.distances {
		first := ps.particles[pair[0]]
		second := ps.particles[pair[1]]
		if first == nil || second == nil {
			safePairs = append(safePairs, pair)
			continue
		}
		newDist := first.Distance(second)
		if newDist >= prevDist && accelerates[pair[0]] && accelerates[pair[1]] {
			safePairs = append(safePairs, pair)
			continue
		}
		ps.distances[pair] = newDist
	}
	for _, pair := range safePairs {
		delete(ps.distances, pair)
	}
}

func (ps *ParticleSystem) String() string {
	return fmt.Sprintf("Particles: %d, Possibly colliding pairs: %d", len(ps.particles), len(ps.distances))
}

func Day20(part int, data []byte) {
	var err error
	lines := strings.Split(string(data), "\n")
	particles := make([]*Particle, len(lines))
	for i, line := range lines {
		particles[i] = new(Particle)
		fields := strings.Split(line, ", ")
		for _, field := range fields {
			vector := strings.Trim(field, "<>=apv")
			scalars := strings.Split(vector, ",")
			if len(scalars) != 3 {
				log.Fatalf("failed parsing vector %s", field)
			}
			var v Vector
			v[0], err = strconv.Atoi(scalars[0])
			if err != nil {
				log.Fatalf("failed parsing number %s at line %d", scalars[0], i)
			}
			v[1], err = strconv.Atoi(scalars[1])
			if err != nil {
				log.Fatalf("failed parsing number %s at line %d", scalars[1], i)
			}
			v[2], err = strconv.Atoi(scalars[2])
			if err != nil {
				log.Fatalf("failed parsing number %s at line %d", scalars[2], i)
			}
			switch field[0] {
			case 'a':
				particles[i].a = v
			case 'v':
				particles[i].v = v
			case 'p':
				particles[i].p = v
			default:
				log.Fatalf("unknown vector %s on line %d", field[0], i)
			}
		}
	}
	log.Printf("Loaded %d particles", len(particles))
	switch part {
	case 1:
		minA := 0.0
		minI := 0
		for i, p := range particles {
			a := p.A()
			if minA == 0 || a < minA {
				minA = a
				minI = i
			}
		}
		log.Printf("In long term closest particle to <0,0,0> is %d and it's acceleration scalar is %f", minI, minA)
	case 2:
		system := NewParticleSystem(particles)
		count := 0
		for system.HasPossiblyColliding() {
			count++
			system.Tick()
		}
		log.Printf("After %d ticks there are %d particles left", count, system.Particles())
	}
}
