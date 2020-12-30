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

	// Build ipTable to reuse data in part2
	ipTable := []ipv7{}

	fileRows := strings.Split(string(fileContent), "\n")
	for _, ip7string := range fileRows {

		ip := convertStringToIP(ip7string)
		ipTable = append(ipTable, ip)
		if abbaSupport(ip) {
			ipCounter++
		}
	}
	fmt.Println("How many IPs in your puzzle input support SSL?")
	fmt.Println(ipCounter)
	// -----------------------------------

	fmt.Println()
	fmt.Println("Part 2: SSL support")
	sslCount := 0
	for _, ip := range ipTable {
		if sslSupport(ip) {
			sslCount++
		}
	}

	fmt.Println("How many IPs in your puzzle input support SSL?")
	fmt.Println(sslCount)
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

// ---------- ABBA for Part 1
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

// ---------- SSL for Part 2
func sslSupport(ip ipv7) bool {
	// Find all abas in support in supernets
	abas := []string{}

	for _, snet := range ip.supernet {
		abas = append(abas, abaSplit(snet)...)
	}

	// ABAs found - check for BABs in hypernets
	for _, aba := range abas {
		bab := aba[1:2] + aba[0:1] + aba[1:2]
		for _, hnet := range ip.hypernet {
			if strings.Index(hnet, bab) > -1 {
				return true
			}
		}
	}

	// No BAB found in hypernet, check failed.
	return false
}

func abaSplit(supernet string) []string {
	abas := []string{}
	for i := 0; i < len(supernet)-2; i++ {
		// ABA, but not AAA
		if supernet[i] == supernet[i+2] && supernet[i] != supernet[i+1] {
			abas = append(abas, supernet[i:i+3])
		}
	}
	return abas
}
