package main

// THIS IS NOT A REAL TEST FILE!
// NEED TO BE REFACTORED INTO REAL TESTSUITE
// THIS WAS JUST USED FOR DELVELOPMENT OF THE main.go

import "fmt"

func main() {
	testPasswords := []string{"hijklmmn", "abbceffg", "abbcegjk", "abcdffaa", "ghjaabcc"}

	for _, testPwd := range testPasswords {
		fmt.Println("Testing:", testPwd)
		if hasIncreasingStraight(testPwd) {
			fmt.Println("pass rule 1")
		} else {
			fmt.Println("failed rule 1")
		}

		if !hasForbiddenLetters(testPwd) {
			fmt.Println("pass rule 2")
		} else {
			fmt.Println("failed rule 2")
		}
		if hasAtLeastTwoPairs(testPwd) {
			fmt.Println("pass rule 3")
		} else {
			fmt.Println("failed rule 3")
		}
	}
}
