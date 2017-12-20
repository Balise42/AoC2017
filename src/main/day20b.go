package main

import (
	"bufio"
	"day20"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("can't open file", err)
	}

	particles := make(map[int]day20.Particle)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		p, v, a := day20.ParseLine(strings.TrimSpace(line))
		particles[i] = day20.Particle{i, p, v, a}
		i = i + 1
	}

	sim := day20.NewSimulation(particles)
	sim.Simulate()
	fmt.Println(len(sim.Particles))
}
