package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type ipv7 struct {
	supernet []string
	hypernet []string
}

func main() {
	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 7: Internet Protocol Version 7")
	fileContent, err := ioutil.ReadFile("puzzle.txt")

	ipCounter := 0

	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileRows := strings.Split(string(fileContent), "\n")
	for _, ip7string := range fileRows {

		ip := convertStringToIP(ip7string)
		if abbaSupport(ip) {
			ipCounter++
		}
	}
	fmt.Println("How many IPs in your puzzle input support TLS?")
	fmt.Println(ipCounter)
	// -----------------------------------

}

func convertStringToIP(s string) ipv7 {
	var ip ipv7
	lastPos := 0
	for i, char := range s {
		if char == '[' {
			ip.supernet = append(ip.supernet, s[lastPos:i])
			lastPos = i + 1
		} else if char == ']' {
			ip.hypernet = append(ip.hypernet, s[lastPos:i])
			lastPos = i + 1
		} else if i == len(s)-1 {
			ip.supernet = append(ip.supernet, s[lastPos:])
		}
	}
	return ip
}

func abbaSupport(ip ipv7) bool {
	for _, hyper := range ip.hypernet {
		if abbaCheck(hyper) {
			// Any abbaCheck true, fails the abbaSupport check instantly
			return false
		}
	}
	// SuperCheck used as just one supernet need to pass.
	superCheck := false
	for _, super := range ip.supernet {
		if abbaCheck(super) {
			superCheck = true
		}
	}
	return superCheck
}

func abbaCheck(s string) bool {
	for x := 0; x < len(s)-3; x++ {
		if s[x] == s[x+3] && s[x+1] == s[x+2] && s[x] != s[x+1] {
			return true
		}
	}

	// abba sequence not found
	return false
}
