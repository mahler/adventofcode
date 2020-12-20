package main

import (
	"fmt"
	"strconv"
)

func main() {
	pwStart := 357253
	pwEnd := 892942

	testSuccess := 0
	for pwTest := pwStart; pwTest < pwEnd; pwTest++ {

		if passwordTest(pwTest, false) {
			testSuccess++
		}
	}
	fmt.Println()
	fmt.Println("Day 04: Secure Container")
	fmt.Println("Possible passwords:", testSuccess)

	fmt.Println()
	fmt.Println("Part 2")

	testSuccess = 0
	for pwTest := pwStart; pwTest < pwEnd; pwTest++ {

		if passwordTest(pwTest, true) {
			testSuccess++
		}
	}
	fmt.Println("Isolated test:", testSuccess)
}

func passwordTest(password int, isoTest bool) bool {
	pass := strconv.FormatInt(int64(password), 10)

	passwordOk := true

	// It is a six-digit number.
	if len(pass) != 6 {
		passwordOk = false
	}

	// Verify that numbers are not decreasing.
	for i := 0; i+1 < len(pass); i++ {
		if pass[i] > pass[i+1] {
			passwordOk = false
		}
	}

	// Verify that there is at least one duplicate pair.
	duplicates := make(map[rune]int, 6)
	duplicatesOk := false
	for _, r := range pass {
		duplicates[r]++
		if duplicates[r] > 1 {
			duplicatesOk = true
		}
	}
	if !duplicatesOk {
		passwordOk = false
	}

	if isoTest {
		seen := make(map[rune]int)
		startVal := rune(0)
		for _, v := range pass {
			if v < startVal {
				return false
			}
			startVal = v
			seen[v] = seen[v] + 1
		}

		double_detected := false
		for _, v := range seen {
			if v == 2 {
				double_detected = true
			}
		}
		if !double_detected {
			passwordOk = false
		}
	}

	return passwordOk
}
