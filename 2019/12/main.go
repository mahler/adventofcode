package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type vector3d struct {
	x, y, z int
}

type moon struct {
	pos vector3d
	vel vector3d
}

// Add method for vector3d to calculate absolute sum
func (v vector3d) absSum() float64 {
	return math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z))
}

func (m *moon) Move() {
	m.pos.x += m.vel.x
	m.pos.y += m.vel.y
	m.pos.z += m.vel.z
}

func (m *moon) Energy() int {
	return int(m.pos.absSum() * m.vel.absSum())
}

// updateVelocity handles velocity changes for a single coordinate
func updateVelocity(pos1, pos2 int) int {
	switch {
	case pos1 < pos2:
		return 1
	case pos1 > pos2:
		return -1
	default:
		return 0
	}
}

func applyGravity(moons []moon) {
	for i := range moons {
		for j := range moons {
			if i == j {
				continue
			}
			moons[i].vel.x += updateVelocity(moons[i].pos.x, moons[j].pos.x)
			moons[i].vel.y += updateVelocity(moons[i].pos.y, moons[j].pos.y)
			moons[i].vel.z += updateVelocity(moons[i].pos.z, moons[j].pos.z)
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / gcd(a, b)
}

// readMoons reads moon data from file and returns a slice of moons
func readMoons(filename string) ([]moon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var moons []moon
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y, z int
		_, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			continue // Skip invalid lines
		}
		moons = append(moons, moon{
			pos: vector3d{x, y, z},
			vel: vector3d{},
		})
	}
	return moons, scanner.Err()
}

// findAxisPeriod finds the period for a single axis
func findAxisPeriod(moons []moon, getVel func(m *moon) int) int {
	moonsCopy := make([]moon, len(moons))
	copy(moonsCopy, moons)

	for i := 1; ; i++ {
		applyGravity(moonsCopy)
		for m := range moonsCopy {
			moonsCopy[m].Move()
		}

		allZeroVel := true
		for m := range moonsCopy {
			if getVel(&moonsCopy[m]) != 0 {
				allZeroVel = false
				break
			}
		}
		if allZeroVel {
			return i * 2
		}
	}
}

func part1(moons []moon) int {
	moonsCopy := make([]moon, len(moons))
	copy(moonsCopy, moons)

	for i := 0; i < 1000; i++ {
		applyGravity(moonsCopy)
		for m := range moonsCopy {
			moonsCopy[m].Move()
		}
	}

	total := 0
	for m := range moonsCopy {
		total += moonsCopy[m].Energy()
	}
	return total
}

func part2(moons []moon) int {
	xPeriod := findAxisPeriod(moons, func(m *moon) int { return m.vel.x })
	yPeriod := findAxisPeriod(moons, func(m *moon) int { return m.vel.y })
	zPeriod := findAxisPeriod(moons, func(m *moon) int { return m.vel.z })

	return lcm(lcm(xPeriod, yPeriod), zPeriod)
}

func main() {
	moons, err := readMoons("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading moons: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Part 1: What is the total energy in the system after simulating the moons given in your scan for 1000 steps?")
	fmt.Println(part1(moons))
	fmt.Println()
	fmt.Println("Part 2: How many steps does it take to reach the first state that exactly matches a previous state?")
	fmt.Println(part2(moons))
}
