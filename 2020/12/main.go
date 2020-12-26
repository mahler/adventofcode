package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type command struct {
	action rune
	value  int
}

type wayPoint struct {
	relNorth int
	relEast  int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	records := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println()
	fmt.Println("Records in Dataset:", len(records))
	fmt.Println()
	commands := make([]command, 0, len(records))

	for _, instruction := range records {
		// Get the action from the instruction.
		a := rune(instruction[0])
		v, _ := strconv.Atoi(string(instruction[1:]))
		commands = append(commands, command{a, v})
	}

	//fmt.Println(commands)
	var north, east, direction int
	for _, command := range commands {
		switch command.action {
		case 'N':
			north += command.value
		case 'E':
			east += command.value
		case 'S':
			north -= command.value
		case 'W':
			east -= command.value
		case 'L':
			direction -= command.value
			for direction < 0 {
				direction = direction + 360
			}
		case 'R':
			direction += command.value
			for direction >= 360 {
				direction = direction - 360
			}
		case 'F':
			switch direction {
			case 0:
				east += command.value
			case 90:
				north -= command.value
			case 180:
				east -= command.value
			case 270:
				north += command.value
			}
		}

	}
	fmt.Println("Day 12, PART 1: Rain Risk")
	fmt.Println("Manhattan distance:", int(math.Abs(float64(north))+math.Abs(float64(east))))

	fmt.Println()
	fmt.Println("Day 12: PART 2")

	waypoint := wayPoint{
		relNorth: 1,
		relEast:  10,
	}

	// Reset north and east
	north, east = 0, 0
	for _, command := range commands {
		switch command.action {
		case 'N':
			waypoint.relNorth += command.value
		case 'S':
			waypoint.relNorth -= command.value
		case 'E':
			waypoint.relEast += command.value
		case 'W':
			waypoint.relEast -= command.value
		case 'L':
			n := int(math.Abs(float64(command.value)) / 90)
			waypoint = rotateWaypoint(false, waypoint, n)
		case 'R':
			n := int(math.Abs(float64(command.value)) / 90)
			waypoint = rotateWaypoint(true, waypoint, n)
		case 'F':
			for i := 0; i < command.value; i++ {
				north += waypoint.relNorth
				east += waypoint.relEast
			}
		}
	}

	fmt.Println("Manhattan distance:", int(math.Abs(float64(north))+math.Abs(float64(east))))
}

func rotateWaypoint(right bool, waypoint wayPoint, n int) wayPoint {
	var newWaypoint wayPoint
	if right {
		newWaypoint.relEast = waypoint.relNorth
		newWaypoint.relNorth = -waypoint.relEast
	} else {
		newWaypoint.relNorth = waypoint.relEast
		newWaypoint.relEast = -waypoint.relNorth
	}
	if n > 1 {
		newWaypoint = rotateWaypoint(right, newWaypoint, n-1)
	}
	return newWaypoint
}
