package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func apply(to *[3]int64, der *[3]int64) {
	to[0] += der[0]
	to[1] += der[1]
	to[2] += der[2]
}

func dist(p *[3]int64) uint64 {
	return uint64(math.Abs(float64(p[0])) + math.Abs(float64(p[1])) + math.Abs(float64(p[2])))
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	re := regexp.MustCompile(`[^\d-]+`)
	var particles [][3][3]int64

	for _, line := range lines {
		if line == "" {
			continue
		}
		nums := re.Split(line[3:], -1)
		var particle [3][3]int64
		for i := 0; i < 9; i++ {
			num, _ := strconv.ParseInt(nums[i], 10, 64)
			particle[i/3][i%3] = num
		}
		particles = append(particles, particle)
	}

	// Part 1
	maxAcc := uint64(0)
	for _, p := range particles {
		accDist := dist(&p[2])
		if accDist > maxAcc {
			maxAcc = accDist
		}
	}

	particlesCopy := make([][3][3]int64, len(particles))
	copy(particlesCopy, particles)

	for i := uint64(0); i < maxAcc*maxAcc; i++ {
		for j := range particlesCopy {
			apply(&particlesCopy[j][1], &particlesCopy[j][2])
			apply(&particlesCopy[j][0], &particlesCopy[j][1])
		}
	}

	minDistIndex := 0
	minDist := dist(&particlesCopy[0][0])
	for i := 1; i < len(particlesCopy); i++ {
		currDist := dist(&particlesCopy[i][0])
		if currDist < minDist {
			minDist = currDist
			minDistIndex = i
		}
	}
	fmt.Println("Part 1: Which particle will stay closest to position <0,0,0> in the long term?")
	fmt.Println(minDistIndex)

	// Part 2
	particles2 := make([][3][3]int64, len(particles))
	copy(particles2, particles)

	for i := uint64(0); i < maxAcc*maxAcc; i++ {
		posCount := make(map[[3]int64]int)
		for j := range particles2 {
			apply(&particles2[j][1], &particles2[j][2])
			apply(&particles2[j][0], &particles2[j][1])
			posCount[particles2[j][0]]++
		}

		var survivingParticles [][3][3]int64
		for _, p := range particles2 {
			if posCount[p[0]] < 2 {
				survivingParticles = append(survivingParticles, p)
			}
		}
		particles2 = survivingParticles
	}

	fmt.Println()
	fmt.Println("Part 2: How many particles are left after all collisions are resolved?")
	fmt.Println(len(particles2))
}
