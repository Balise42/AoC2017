package day20

import (
	"fmt"
	"testing"
)

func TestEquality(t *testing.T) {
	p1 := Vector{-13, 1, 99}
	p2 := Vector{-13, 1, 99}
	if p1 != p2 {
		t.Error("not equal")
	}
}

func TestSimulation(t *testing.T) {
	line1 := "p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>"
	line2 := "p=<4,0,0>, v=<0,0,0>, a=<0,0,0>"
	line3 := "p=<56,0,0>, v=<1,1,1>, a=<8,8,8>"
	particles := make(map[int]Particle)
	p1a, p1v, p1p := ParseLine(line1)
	particles[0] = Particle{0, p1p, p1v, p1a}
	p2a, p2v, p2p := ParseLine(line2)
	particles[1] = Particle{1, p2p, p2v, p2a}
	p3a, p3v, p3p := ParseLine(line3)
	particles[2] = Particle{2, p3p, p3v, p3a}

	sim := NewSimulation(particles)
	sim.Simulate()
	fmt.Println(sim)
}
