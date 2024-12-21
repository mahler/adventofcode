package main

import (
	"fmt"
	"sort"
)

var startFloors = [][]string{
	{"OG", "TG", "TM", "PG", "RG", "RM", "CG", "CM"},
	{"OM", "PM"},
	{},
	{},
}

func checkValid(dest int, pt1 string, pt2 string, validFloors [][]string, validElevator int) bool {
	testDest := make([]string, len(validFloors[dest]))
	copy(testDest, validFloors[dest])
	testDest = append(testDest, pt1)
	if pt2 != "" {
		testDest = append(testDest, pt2)
	}

	for _, i := range testDest {
		if i[1] == 'M' {
			for _, j := range testDest {
				if j[1] == 'G' && !contains(testDest, string(i[0])+"G") {
					return false
				}
			}
		}
	}

	if !contains(validFloors[validElevator], pt1) {
		return false
	}
	if pt2 != "" && !contains(validFloors[validElevator], pt2) {
		return false
	}

	testFloor := make([]string, len(validFloors[validElevator]))
	copy(testFloor, validFloors[validElevator])
	testFloor = remove(testFloor, pt1)
	if pt2 != "" {
		testFloor = remove(testFloor, pt2)
	}

	for _, i := range testFloor {
		if i[1] == 'M' && !contains(testFloor, string(i[0])+"G") {
			for _, j := range testFloor {
				if j != i && j[1] == 'G' {
					return false
				}
			}
		}
	}
	return true
}

type State [5][3]int

type MoveResult struct {
	floors   [][]string
	elevator int
	state    State
	steps    int
}

func doMove(destination int, part1 string, part2 string, inFloors [][]string, inElevator int, steps int) *MoveResult {
	moveFloors := make([][]string, len(inFloors))
	for i := range inFloors {
		moveFloors[i] = make([]string, len(inFloors[i]))
		copy(moveFloors[i], inFloors[i])
	}
	moveElevator := inElevator

	valid := checkValid(destination, part1, part2, moveFloors, moveElevator)
	state := State{}
	state[4][0] = destination

	if !valid {
		return nil
	}

	moveFloors[destination] = append(moveFloors[destination], part1)
	if part2 != "" {
		moveFloors[destination] = append(moveFloors[destination], part2)
	}
	moveFloors[moveElevator] = remove(moveFloors[moveElevator], part1)
	if part2 != "" {
		moveFloors[moveElevator] = remove(moveFloors[moveElevator], part2)
	}
	sort.Strings(moveFloors[destination])
	moveElevator = destination

	for i := 0; i < 4; i++ {
		for _, j := range moveFloors[i] {
			if j[1] == 'M' {
				if !contains(moveFloors[i], string(j[0])+"G") {
					state[i][0]++
				} else {
					state[i][2]++
				}
			} else {
				if !contains(moveFloors[i], string(j[0])+"M") {
					state[i][1]++
				}
			}
		}
	}

	return &MoveResult{
		floors:   moveFloors,
		elevator: moveElevator,
		state:    state,
		steps:    steps,
	}
}

func findPath(startFloors [][]string, amount int) int {
	steps := 1
	startElevator := 0
	moves := []*MoveResult{}
	seen := map[State]bool{}
	found := false

	for _, i := range startFloors[startElevator] {
		test := doMove(1, i, "", startFloors, startElevator, steps)
		if test != nil && !seen[test.state] {
			moves = append(moves, test)
			seen[test.state] = true
		}
		for _, j := range startFloors[startElevator] {
			if i != j {
				test := doMove(1, i, j, startFloors, startElevator, steps)
				if test != nil && !seen[test.state] {
					moves = append(moves, test)
					seen[test.state] = true
				}
			}
		}
	}

	for !found {
		steps++
		currentMoves := make([]*MoveResult, len(moves))
		copy(currentMoves, moves)
		moves = []*MoveResult{}

		for _, i := range currentMoves {
			if len(i.floors[3]) == amount {
				return i.steps
			}

			testFloors := make([][]string, len(i.floors))
			for f := range i.floors {
				testFloors[f] = make([]string, len(i.floors[f]))
				copy(testFloors[f], i.floors[f])
			}
			testElevator := i.elevator

			for _, p1 := range testFloors[testElevator] {
				for _, p2 := range testFloors[testElevator] {
					if p1 != p2 {
						if testElevator < 3 {
							test := doMove(testElevator+1, p1, p2, testFloors, testElevator, steps)
							if test != nil && !seen[test.state] {
								moves = append(moves, test)
								seen[test.state] = true
							}
						}
					}
				}

				if testElevator < 3 {
					test := doMove(testElevator+1, p1, "", testFloors, testElevator, steps)
					if test != nil && !seen[test.state] {
						moves = append(moves, test)
						seen[test.state] = true
					}
				}

				if testElevator > 0 {
					allEmpty := true
					for x := 1; x < testElevator; x++ {
						if len(testFloors[x]) > 0 {
							allEmpty = false
							break
						}
					}
					if !allEmpty || testElevator <= 1 {
						test := doMove(testElevator-1, p1, "", testFloors, testElevator, steps)
						if test != nil && !seen[test.state] {
							moves = append(moves, test)
							seen[test.state] = true
						}

						for _, p2 := range testFloors[testElevator] {
							if p1[0] != p2[0] {
								test := doMove(testElevator-1, p1, p2, testFloors, testElevator, steps)
								if test != nil && !seen[test.state] {
									moves = append(moves, test)
									seen[test.state] = true
								}
							}
						}
					}
				}
			}
		}
	}
	return -1
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	floors1 := [][]string{
		{"OG", "TG", "TM", "PG", "RG", "RM", "CG", "CM"},
		{"OM", "PM"},
		{},
		{},
	}
	fmt.Println("Part 1: In your situation, what is the minimum number of steps required to bring all of the objects to the fourth floor?")
	fmt.Println(findPath(floors1, 10))

	floors2 := [][]string{
		{"EG", "EM", "DG", "DM", "OG", "TG", "TM", "PG", "RG", "RM", "CG", "CM"},
		{"OM", "PM"},
		{},
		{},
	}
	fmt.Println()
	fmt.Println("Part 2: What is the minimum number of steps required to bring all of the objects, including these four new ones, to the fourth floor?")
	fmt.Println(findPath(floors2, 14))
}
