package day20

import "fmt"

type Pair struct {
	A int
	B int
}

type Particle struct {
	Id int
	P  Vector
	V  Vector
	A  Vector
}

type Delta struct {
	DeltaP int
	DeltaV int
}

type Simulation struct {
	Particles  map[int]Particle
	Candidates map[Pair]Delta
}

func NewSimulation(particles map[int]Particle) Simulation {
	candidates := make(map[Pair]Delta)
	for i, part1 := range particles {
		for j := i + 1; j < len(particles); j++ {
			part2 := particles[j]
			deltaP := dist(part1.P, part2.P)
			deltaV := dist(part1.V, part2.V)
			candidates[Pair{i, j}] = Delta{deltaP, deltaV}
		}
	}
	sim := Simulation{particles, candidates}
	return sim
}

func dist(a Vector, b Vector) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z)
}

func (sim *Simulation) Simulate() {
	for len(sim.Candidates) > 0 {
		fmt.Println(len(sim.Particles))
		fmt.Println(len(sim.Candidates))
		newState := make(map[int]Particle)
		for key, val := range sim.Particles {
			a := Vector(val.A)
			v := Vector{val.V.X + a.X, val.A.Y + a.Y, val.A.Z + a.Z}
			p := Vector{val.P.X + v.X, val.P.Y + v.Y, val.P.Z + v.Z}
			newState[key] = Particle{key, p, v, a}
		}
		for p, delta := range sim.Candidates {
			if newState[p.A].P == newState[p.B].P {
				delete(newState, p.A)
				delete(newState, p.B)
				for k := range sim.Candidates {
					fmt.Println(k, p)
					if k.A == p.A || k.B == p.A || k.A == p.B || k.B == p.B {
						delete(sim.Candidates, k)
					}
				}
			}
			newDeltaP := dist(newState[p.A].P, newState[p.B].P)
			newDeltaV := dist(newState[p.A].V, newState[p.B].V)
			if delta.DeltaP <= newDeltaP && delta.DeltaV <= newDeltaV {
				delete(sim.Candidates, p)
			} else {
				if _, ok := sim.Candidates[p]; ok {
					sim.Candidates[p] = Delta{newDeltaP, newDeltaV}
				}
			}
		}
		sim.Particles = newState
	}
}
