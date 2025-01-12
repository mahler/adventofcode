package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func propagatePulse(graph map[string][]string, flops map[string]bool, conjs map[string]map[string]bool, sender, receiver string, pulse bool) []struct {
	sender, receiver string
	pulse            bool
} {
	var nextPulse bool
	result := []struct {
		sender, receiver string
		pulse            bool
	}{}

	if _, ok := flops[receiver]; ok {
		if pulse {
			return result
		}
		nextPulse = !flops[receiver]
		flops[receiver] = nextPulse
	} else if _, ok := conjs[receiver]; ok {
		conjs[receiver][sender] = pulse
		nextPulse = true
		for _, val := range conjs[receiver] {
			if !val {
				nextPulse = false
				break
			}
		}
		nextPulse = !nextPulse
	} else if _, ok := graph[receiver]; ok {
		nextPulse = pulse
	} else {
		return result
	}

	for _, newReceiver := range graph[receiver] {
		result = append(result, struct {
			sender, receiver string
			pulse            bool
		}{receiver, newReceiver, nextPulse})
	}

	return result
}

func run(graph map[string][]string, flops map[string]bool, conjs map[string]map[string]bool) (int, int) {
	queue := list.New()
	queue.PushBack(struct {
		sender, receiver string
		pulse            bool
	}{"button", "broadcaster", false})

	nhi, nlo := 0, 0

	for queue.Len() > 0 {
		elem := queue.Remove(queue.Front()).(struct {
			sender, receiver string
			pulse            bool
		})

		sender, receiver, pulse := elem.sender, elem.receiver, elem.pulse
		if pulse {
			nhi++
		} else {
			nlo++
		}

		for _, newElem := range propagatePulse(graph, flops, conjs, sender, receiver, pulse) {
			queue.PushBack(newElem)
		}
	}

	return nhi, nlo
}

func findPeriods(graph map[string][]string, flops map[string]bool, conjs map[string]map[string]bool) []int {
	periodic := make(map[string]bool)
	var rxSource string

	for source, dests := range graph {
		if len(dests) == 1 && dests[0] == "rx" {
			if _, ok := conjs[source]; ok {
				rxSource = source
				break
			}
		}
	}

	for source, dests := range graph {
		for _, dest := range dests {
			if dest == rxSource {
				if _, ok := conjs[source]; ok {
					periodic[source] = true
				}
			}
		}
	}

	var iterations []int
	for iteration := 1; ; iteration++ {
		queue := list.New()
		queue.PushBack(struct {
			sender, receiver string
			pulse            bool
		}{"button", "broadcaster", false})

		for queue.Len() > 0 {
			elem := queue.Remove(queue.Front()).(struct {
				sender, receiver string
				pulse            bool
			})

			sender, receiver, pulse := elem.sender, elem.receiver, elem.pulse

			if !pulse {
				if _, ok := periodic[receiver]; ok {
					iterations = append(iterations, iteration)
					delete(periodic, receiver)
					if len(periodic) == 0 {
						return iterations
					}
				}
			}

			for _, newElem := range propagatePulse(graph, flops, conjs, sender, receiver, pulse) {
				queue.PushBack(newElem)
			}
		}
	}
}

func lcm(numbers []int) int {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = result * num / gcd(result, num)
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	flops := make(map[string]bool)
	conjs := make(map[string]map[string]bool)
	graph := make(map[string][]string)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")
		source := strings.TrimSpace(parts[0])
		dests := strings.Split(strings.TrimSpace(parts[1]), ", ")

		if source[0] == '%' {
			source = source[1:]
			flops[source] = false
		} else if source[0] == '&' {
			source = source[1:]
			conjs[source] = make(map[string]bool)
		}

		graph[source] = dests
	}

	for source, dests := range graph {
		for _, dest := range dests {
			if _, ok := conjs[dest]; ok {
				conjs[dest][source] = false
			}
		}
	}

	tothi, totlo := 0, 0
	for i := 0; i < 1000; i++ {
		nhi, nlo := run(graph, flops, conjs)
		tothi += nhi
		totlo += nlo
	}

	part1solution := totlo * tothi
	fmt.Println("Part 1: What do you get if you multiply the total number of low pulses sent by the total number of high pulses sent?")
	fmt.Println(part1solution)

	// Part 2
	for f := range flops {
		flops[f] = false
	}

	for _, inputs := range conjs {
		for k := range inputs {
			inputs[k] = false
		}
	}

	periods := findPeriods(graph, flops, conjs)
	part2solution := lcm(periods)
	fmt.Println("Part 2: Reset all modules to their default states. Waiting for all pulses to be fully handled after each button press,")
	fmt.Println("what is the fewest number of button presses required to deliver a single low pulse to the module named rx?")
	fmt.Println(part2solution)
}
