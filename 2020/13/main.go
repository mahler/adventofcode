package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("DAY13, Part 1: Shuttle Search")

	data, err := os.ReadFIle("shuttle.schedule")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	records := strings.Split(string(data), "\n")

	// first line has timestamp.
	timestamp, err := strconv.ParseInt(records[0], 10, 64)

	busEntries := strings.Split(records[1], ",")
	busses := []int64{}

	for _, bus := range busEntries {
		if bus == "x" {
			continue
		}
		val, _ := strconv.ParseInt(bus, 10, 64)
		busses = append(busses, val)
	}

	//result := solve(timestamp, busses)
	var minInterval int64 = 1<<63 - 1
	var minBusId int64

	for _, busId := range busses {
		lastTime := timestamp % busId
		nextTime := busId - lastTime
		if nextTime < minInterval {
			minInterval = nextTime
			minBusId = busId
		}
	}
	fmt.Println("Result for part 1:", minBusId*minInterval)

	fmt.Println()
	fmt.Println("DAY13, Part 2: Shuttle Search")

	p2busses := make(map[int64]int64)

	for i, bus := range busEntries {
		if bus == "x" {
			continue
		}
		val, _ := strconv.ParseInt(bus, 10, 64)
		if i == 0 {
			p2busses[val] = 0
		} else {
			p2busses[val] = val - int64(i)%val
		}
	}

	var skips int64
	var offset int64

	keys := make([]int64, 0, len(p2busses))
	for k := range p2busses {
		keys = append(keys, k)
	}

	// For any group, the pattern repeats every LCM of the elements involved.
	// The offset is the offset from that LCM at which the pattern starts.
	// That offset becomes the new remainder for the next step.
	for _, k := range keys {
		v := p2busses[k]
		if skips == 0 {
			skips = k
			offset = v
			continue
		}
		lcm := Lcm(skips, k)
		for i := offset % lcm; i <= lcm; i += skips {
			if i%k == v {
				skips = lcm
				offset = i
				break
			}
		}
	}

	fmt.Println("Earliest timestamp :", offset)

}

// Greatest Common Denominator
func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Least Common Multiple - https://en.wikipedia.org/wiki/Least_common_multiple
func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}
