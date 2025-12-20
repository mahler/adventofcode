package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func rep1(n int) []int {
	result := make([]int, 0, 9)
	for c := 1; c <= 9; c++ {
		sum := 0
		for i := 0; i < n; i++ {
			sum += c * int(math.Pow10(i))
		}
		result = append(result, sum)
	}
	return result
}

func buildRep() map[int]map[int]bool {
	rep := make(map[int]map[int]bool)

	rep[1] = make(map[int]bool)

	rep[2] = make(map[int]bool)
	for c := 1; c <= 9; c++ {
		rep[2][c*10+c] = true
	}

	rep[3] = make(map[int]bool)
	for _, v := range rep1(3) {
		rep[3][v] = true
	}

	rep[4] = make(map[int]bool)
	for r := 10; r < 100; r++ {
		rep[4][r*100+r] = true
	}

	rep[5] = make(map[int]bool)
	for _, v := range rep1(5) {
		rep[5][v] = true
	}

	rep[6] = make(map[int]bool)
	for r := 10; r < 100; r++ {
		rep[6][r*10000+r*100+r] = true
	}
	for r := 100; r < 1000; r++ {
		rep[6][r*1000+r] = true
	}

	rep[7] = make(map[int]bool)
	for _, v := range rep1(7) {
		rep[7][v] = true
	}

	rep[8] = make(map[int]bool)
	for r := 1000; r < 10000; r++ {
		rep[8][r*10000+r] = true
	}

	rep[9] = make(map[int]bool)
	for r := 100; r < 1000; r++ {
		rep[9][r*1000000+r*1000+r] = true
	}

	rep[10] = make(map[int]bool)
	for r := 10; r < 100; r++ {
		rep[10][r*100000000+r*1000000+r*10000+r*100+r] = true
	}
	for r := 10000; r < 100000; r++ {
		rep[10][r*100000+r] = true
	}

	return rep
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}
	return len(strconv.Itoa(n))
}

func checkRange(s string, rep map[int]map[int]bool) ([]int, []int) {
	parts := strings.Split(s, "-")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])

	// Union of rep sets
	rs := make(map[int]bool)
	for k := range rep[numDigits(a)] {
		rs[k] = true
	}
	for k := range rep[numDigits(b)] {
		rs[k] = true
	}

	// Filter to range [a, b]
	all := make([]int, 0)
	for r := range rs {
		if r >= a && r <= b {
			all = append(all, r)
		}
	}

	// Filter even (palindrome check)
	even := make([]int, 0)
	for _, r := range all {
		digits := numDigits(r)
		h := int(math.Pow10(digits / 2))
		if r/h == r%h {
			even = append(even, r)
		}
	}

	return even, all
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	rep := buildRep()

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")

	p1, p2 := 0, 0
	for _, r := range ranges {
		even, all := checkRange(r, rep)
		p1 += sum(even)
		p2 += sum(all)
	}
	fmt.Println("What do you get if you add up all of the invalid IDs?")
	fmt.Println()
	fmt.Println(p1)
	fmt.Println("What do you get if you add up all of the invalid IDs using these new rules?")
	fmt.Println(p2)
}
