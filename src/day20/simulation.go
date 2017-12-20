package day20

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
		newState := make(map[int]Particle)
		for key, val := range sim.Particles {
			a := Vector(val.A)
			v := Vector{val.V.X + a.X, val.V.Y + a.Y, val.V.Z + a.Z}
			p := Vector{val.P.X + v.X, val.P.Y + v.Y, val.P.Z + v.Z}
			newState[key] = Particle{key, p, v, a}
		}
		toDelete := computeColliding(newState)
		for k := range toDelete {
			delete(newState, k)
			for c := range sim.Candidates {
				if c.A == k || c.B == k {
					delete(sim.Candidates, c)
				}
			}
		}
		sim.Particles = newState
		deleteCandidates(sim)
	}
}

func deleteCandidates(sim *Simulation) {
	for k, v := range sim.Candidates {
		if dist(sim.Particles[k.A].P, sim.Particles[k.B].P) >= v.DeltaP && dist(sim.Particles[k.A].V, sim.Particles[k.B].V) >= v.DeltaV {
			delete(sim.Candidates, k)
		} else {
			sim.Candidates[k] = Delta{dist(sim.Particles[k.A].P, sim.Particles[k.B].P), dist(sim.Particles[k.A].V, sim.Particles[k.B].V)}
		}
	}
}

func computeColliding(particles map[int]Particle) map[int]bool {
	toDelete := make(map[int]bool)
	for k1, v1 := range particles {
		for k2, v2 := range particles {
			if k1 == k2 {
				continue
			}
			if v1.P == v2.P {
				toDelete[k1] = true
				toDelete[k2] = true
			}
		}
	}
	return toDelete
}
