package main

import "fmt"

func main() {
	input := "vzbxkghb"

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 11: Corporate Policy")
	for {
		input = nextPassword(input)
		if hasIncreasingStraight(input) && !hasForbiddenLetters(input) && hasAtLeastTwoPairs(input) {
			fmt.Println("Santa's new password:", input)
			break
		}
	}
}

func nextPassword(password string) string {
	pwd := []byte(password)
	for i := 7; i >= 0; i-- {
		pwd[i]++
		if pwd[i] > 122 {
			pwd[i] -= 26
		} else {
			break
		}
	}

	return string(pwd)
}

// Rule 1) Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
func hasIncreasingStraight(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		c := byte(password[i])
		c2 := byte(password[i+1])
		c3 := byte(password[i+2])

		if c+1 == c2 && c2+1 == c3 {
			return true
		}
	}
	// Didn't return earlier, so no increasing string.
	return false
}

// Rule 2) Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
func hasForbiddenLetters(password string) bool {
	for i := 0; i < len(password); i++ {
		if password[i:i+1] == "i" || password[i:i+1] == "o" || password[i:i+1] == "l" {
			return true
		}
	}
	return false
}

// Rule 3) must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
func hasAtLeastTwoPairs(password string) bool {
	leftoverStr := ""
	firstPair := ""

	for x := 0; x < len(password)-1; x++ {
		if password[x:x+1] == password[x+1:x+2] {
			// Found first pair
			firstPair = password[x : x+2]
			leftoverStr = password[x+2:]
			break
		}
	}
	if len(leftoverStr) < 2 {
		// cannot have a second pair.
		return false
	}

	for x := 0; x < len(leftoverStr)-1; x++ {
		if leftoverStr[x:x+1] == leftoverStr[x+1:x+2] && leftoverStr[x:x+2] != firstPair {
			// Found a second pair which is differnt.
			return true
		}
	}
	// Did not fullfil the tests
	return false
}
