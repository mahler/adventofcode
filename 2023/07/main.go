package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strength(hand string, part2 bool) (int, string) {
	// Replace card ranks
	hand = strings.ReplaceAll(hand, "T", string('9'+1))
	if part2 {
		hand = strings.ReplaceAll(hand, "J", string('2'-1))
	} else {
		hand = strings.ReplaceAll(hand, "J", string('9'+2))
	}
	hand = strings.ReplaceAll(hand, "Q", string('9'+3))
	hand = strings.ReplaceAll(hand, "K", string('9'+4))
	hand = strings.ReplaceAll(hand, "A", string('9'+5))

	// Count card frequencies
	counter := make(map[rune]int)
	for _, card := range hand {
		counter[card]++
	}

	if part2 {
		var target rune
		for k := range counter {
			if k != '1' && (counter[k] > counter[target] || target == '1') {
				target = k
			}
		}
		if counter['1'] > 0 && target != '1' {
			counter[target] += counter['1']
			delete(counter, '1')
		}
	}

	// Collect frequencies and sort them
	values := []int{}
	for _, v := range counter {
		values = append(values, v)
	}
	sort.Ints(values)

	// Determine strength
	switch {
	case len(values) == 1 && values[0] == 5:
		return 10, hand
	case len(values) == 2 && values[0] == 1 && values[1] == 4:
		return 9, hand
	case len(values) == 2 && values[0] == 2 && values[1] == 3:
		return 8, hand
	case len(values) == 3 && values[0] == 1 && values[1] == 1 && values[2] == 3:
		return 7, hand
	case len(values) == 3 && values[0] == 1 && values[1] == 2 && values[2] == 2:
		return 6, hand
	case len(values) == 4 && values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 2:
		return 5, hand
	case len(values) == 5 && values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 1 && values[4] == 1:
		return 4, hand
	default:
		panic(fmt.Sprintf("Unexpected case: %v %s %v", counter, hand, values))
	}
}

func main() {
	// Open and read input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, part2 := range []bool{false, true} {
		hands := []struct {
			hand string
			bid  int
		}{}

		for _, line := range lines {
			parts := strings.Fields(line)
			bid, _ := strconv.Atoi(parts[1])
			hands = append(hands, struct {
				hand string
				bid  int
			}{hand: parts[0], bid: bid})
		}

		// Sort hands by strength
		sort.Slice(hands, func(i, j int) bool {
			si, hi := strength(hands[i].hand, part2)
			sj, hj := strength(hands[j].hand, part2)
			if si != sj {
				return si < sj
			}
			return hi < hj
		})

		// Calculate the answer
		ans := 0
		for i, h := range hands {
			ans += (i + 1) * h.bid
		}
		if part2 {
			fmt.Println("Part2")
		} else {
			fmt.Println("Part1")
		}
		fmt.Println(ans)
	}
}
