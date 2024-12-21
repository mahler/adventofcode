package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pwd string = "abcdefgh"

func swapPos(str string, x, y int) string {
	runes := []rune(str)
	runes[x], runes[y] = runes[y], runes[x]
	return string(runes)
}

func swapLetter(str string, x, y rune) string {
	result := strings.ReplaceAll(str, string(x), "&")
	result = strings.ReplaceAll(result, string(y), string(x))
	return strings.ReplaceAll(result, "&", string(y))
}

func rotate(str string, dirstr string, count int) string {
	runes := []rune(str)
	length := len(runes)
	count = count % length

	if dirstr == "left" {
		return string(append(runes[count:], runes[:count]...))
	} else {
		return string(append(runes[length-count:], runes[:length-count]...))
	}
}

func rotateIndex(str string, letter rune) string {
	idx := strings.IndexRune(str, letter)
	rotations := 1 + idx
	if idx >= 4 {
		rotations++
	}
	return rotate(str, "right", rotations)
}

func reverse(str string, istart, iend int) string {
	runes := []rune(str)
	for i := 0; i < (iend-istart+1)/2; i++ {
		runes[istart+i], runes[iend-i] = runes[iend-i], runes[istart+i]
	}
	return string(runes)
}

func move(str string, ia, ib int) string {
	runes := []rune(str)
	char := runes[ia]
	// Remove character at ia
	runes = append(runes[:ia], runes[ia+1:]...)
	// Insert at ib
	runes = append(runes[:ib], append([]rune{char}, runes[ib:]...)...)
	return string(runes)
}

func parseLine(line string) {
	tokens := strings.Fields(line)

	if strings.HasPrefix(line, "swap position") {
		x, _ := strconv.Atoi(tokens[2])
		y, _ := strconv.Atoi(tokens[5])
		pwd = swapPos(pwd, x, y)
	} else if strings.HasPrefix(line, "swap letter") {
		pwd = swapLetter(pwd, rune(tokens[2][0]), rune(tokens[5][0]))
	} else if strings.HasPrefix(line, "rotate based") {
		pwd = rotateIndex(pwd, rune(tokens[6][0]))
	} else if strings.HasPrefix(line, "rotate") {
		count, _ := strconv.Atoi(tokens[2])
		pwd = rotate(pwd, tokens[1], count)
	} else if strings.HasPrefix(line, "reverse") {
		start, _ := strconv.Atoi(tokens[2])
		end, _ := strconv.Atoi(tokens[4])
		pwd = reverse(pwd, start, end)
	} else if strings.HasPrefix(line, "move") {
		a, _ := strconv.Atoi(tokens[2])
		b, _ := strconv.Atoi(tokens[5])
		pwd = move(pwd, a, b)
	} else {
		fmt.Printf("!! %s\n", line)
	}
}

func parseLine2(line, inp string) string {
	tokens := strings.Fields(line)

	if strings.HasPrefix(line, "swap position") {
		x, _ := strconv.Atoi(tokens[2])
		y, _ := strconv.Atoi(tokens[5])
		return swapPos(inp, x, y)
	} else if strings.HasPrefix(line, "swap letter") {
		return swapLetter(inp, rune(tokens[2][0]), rune(tokens[5][0]))
	} else if strings.HasPrefix(line, "rotate based") {
		return rotateIndex(inp, rune(tokens[6][0]))
	} else if strings.HasPrefix(line, "rotate") {
		count, _ := strconv.Atoi(tokens[2])
		return rotate(inp, tokens[1], count)
	} else if strings.HasPrefix(line, "reverse") {
		start, _ := strconv.Atoi(tokens[2])
		end, _ := strconv.Atoi(tokens[4])
		return reverse(inp, start, end)
	} else if strings.HasPrefix(line, "move") {
		a, _ := strconv.Atoi(tokens[2])
		b, _ := strconv.Atoi(tokens[5])
		return move(inp, a, b)
	}
	fmt.Printf("!! %s\n", line)
	return inp
}

// Generate all permutations of a string
func generatePermutations2(str string) []string {
	var result []string
	if len(str) <= 1 {
		return []string{str}
	}

	for i, ch := range str {
		// Remove current character
		remaining := str[:i] + str[i+1:]

		// Generate permutations of remaining characters
		subPerms := generatePermutations2(remaining)

		// Add current character to beginning of each sub-permutation
		for _, perm := range subPerms {
			result = append(result, string(ch)+perm)
		}
	}
	return result
}

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parseLine(line)
		//fmt.Println(pwd)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1: Given the list of scrambling operations in your puzzle input, what is the result of scrambling abcdefgh?	")
	fmt.Println(pwd)

	// --------- Part 2 -------
	pwd := "abcdefgh"

	// Read input file
	file, err = os.Open("puzzle.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var unscrambledPassword string
	// Generate and test all permutations
	perms := generatePermutations2(pwd)
	for _, perm := range perms {
		currentPwd := perm
		for _, line := range lines {
			currentPwd = parseLine2(line, currentPwd)
		}
		if currentPwd == "fbgdceah" {
			unscrambledPassword = perm
			break
		}
	}

	fmt.Println()
	fmt.Println("Part 2: What is the un-scrambled version of the scrambled password fbgdceah?")
	fmt.Println(unscrambledPassword)
}
