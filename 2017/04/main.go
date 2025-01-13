package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	passPhrases := parseInput(string(input))

	valid := countValid(passPhrases, validate)
	secureValid := countValid(passPhrases, validateSecure)

	fmt.Println("Part 1: The system's full passphrase list is available as your puzzle input. How many passphrases are valid?")
	fmt.Println(valid)

	fmt.Println()
	fmt.Println("Part 2: Under this new system policy, how many passphrases are valid?")
	fmt.Println(secureValid)
}

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]string, len(lines))
	for i, line := range lines {
		result[i] = parsePassphrase(line)
	}
	return result
}

func parsePassphrase(line string) []string {
	words := strings.Fields(line)
	result := make([]string, 0, len(words))
	for _, word := range words {
		if word = strings.TrimSpace(word); word != "" {
			result = append(result, word)
		}
	}
	return result
}

func validate(passphrase []string) bool {
	seen := make(map[string]bool)
	for _, word := range passphrase {
		if seen[word] {
			return false
		}
		seen[word] = true
	}
	return true
}

func validateSecure(passphrase []string) bool {
	seen := make(map[string]bool)
	for _, word := range passphrase {
		// Convert to runes, sort, and convert back to string
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool { return runes[i] > runes[j] })
		sorted := string(runes)

		if seen[sorted] {
			return false
		}
		seen[sorted] = true
	}
	return true
}

func countValid(passPhrases [][]string, validationFunc func([]string) bool) int {
	count := 0
	for _, phrase := range passPhrases {
		if validationFunc(phrase) {
			count++
		}
	}
	return count
}
