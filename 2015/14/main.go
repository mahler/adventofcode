package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type reindeer struct {
	name         string
	speed        int
	speedSeconds int
	restSeconds  int
	// Part2 extension
	distance int
	segment  int
	points   int
}

func main() {
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	// https://regoio.herokuapp.com/
	reindeerRX := regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	fmt.Println("Instructions in dataset:", len(lines))
	fmt.Println("lines", len(lines))

	santasReindeer := []reindeer{}
	maxDistance := 0
	maxName := ""

	for _, line := range lines {
		var r reindeer
		fields := reindeerRX.FindStringSubmatch(line)
		r.name = fields[1]
		r.speed, _ = strconv.Atoi(fields[2])
		r.speedSeconds, _ = strconv.Atoi(fields[3])
		r.restSeconds, _ = strconv.Atoi(fields[4])

		//fmt.Println("Name:", r.name, "speed:", r.speed, "speedtime:", r.speedSeconds, "rest:", r.restSeconds)
		santasReindeer = append(santasReindeer, r)
		reinDistance := r.distanceBySecondes(2503)

		if reinDistance > maxDistance {
			maxDistance = reinDistance
			maxName = r.name
		}
	}
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 14, part 1: Reindeer Olympics")
	fmt.Println(maxName, "won with a distance of", maxDistance)

	// Part 2 -----------------------------------
	fmt.Println()
	fmt.Println("Part 2")
	rDeer, points := runReindeerRace(santasReindeer, 2503)

	fmt.Println("Winner is", rDeer.name, "with total points of", points)
}

// distanceBySconds will calculate how far a reindeer will get in secondsToMove
func (r reindeer) distanceBySecondes(secondsToMove int) int {
	distance := 0
	segment := 1
	for x := 0; x <= secondsToMove; x++ {
		//fmt.Println(x, ": segment", segment, "distance:", distance)
		if segment > 0 {
			distance += r.speed
		}
		if segment == r.speedSeconds {
			segment = 0 - r.restSeconds
		}
		segment++
	}
	return distance
}

func runReindeerRace(reindeers []reindeer, raceTime int) (reindeer, int) {

	// For every second of the race
	for x := 0; x < raceTime; x++ {
		maxDistance := 0
		for rDeer := range reindeers {
			thisDeer := reindeers[rDeer]
			// ---
			//fmt.Println(x, ": segment", segment, "distance:", distance)
			if thisDeer.segment > 0 {
				thisDeer.distance += thisDeer.speed
			}
			if thisDeer.segment == thisDeer.speedSeconds {
				thisDeer.segment = 0 - thisDeer.restSeconds
			}
			thisDeer.segment++

			if thisDeer.distance > maxDistance {
				maxDistance = thisDeer.distance
			}
			// ---
			reindeers[rDeer] = thisDeer
		}
		// Finished moving reindeers, assign point for leader(s).
		for rDeer := range reindeers {
			if reindeers[rDeer].distance == maxDistance {
				reindeers[rDeer].points++
			}
		}
	}

	roundLead := 0
	maxPoints := 0
	for rDeer := range reindeers {
		if reindeers[rDeer].points > maxPoints {
			maxPoints = reindeers[rDeer].points
			roundLead = rDeer
		}
	}

	// One round too many of round points, so subtract one.
	maxPoints--

	return reindeers[roundLead], maxPoints
}
