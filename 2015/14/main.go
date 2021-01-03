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
