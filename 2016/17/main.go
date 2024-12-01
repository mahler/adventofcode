package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var data = "awrkjxxr"
var frontier = []struct {
	path     string
	position [2]int
}{}

var pos = [2]int{0, 0}
var vault = [2]int{3, 3}
var found = false
var maxlen = 0

func neighbors(node string) []string {
	doors := []string{"u", "d", "l", "r"}
	hash := md5.New()
	hash.Write([]byte(node))
	hsh := fmt.Sprintf("%x", hash.Sum(nil))
	var unlocked []string
	for i, c := range hsh[:4] {
		if strings.Contains("bcdef", string(c)) {
			unlocked = append(unlocked, doors[i])
		}
	}
	return unlocked
}

func main() {
	frontier = append(frontier, struct {
		path     string
		position [2]int
	}{data, pos})

	fmt.Println("Part 1: Given your vault's passcode, what is the shortest path to reach the vault?")
	for len(frontier) > 0 {
		// Pop the first element
		working := frontier[0].path
		position := frontier[0].position
		frontier = frontier[1:]

		if position == vault {
			if !found {
				fmt.Println(working[8:])
				found = true
			}
			if len(working[8:]) > maxlen {
				maxlen = len(working[8:])
			}
			continue
		}

		for _, c := range neighbors(working) {
			if c == "u" && position[1] > 0 {
				frontier = append(frontier, struct {
					path     string
					position [2]int
				}{working + "U", [2]int{position[0], position[1] - 1}})
			} else if c == "d" && position[1] < 3 {
				frontier = append(frontier, struct {
					path     string
					position [2]int
				}{working + "D", [2]int{position[0], position[1] + 1}})
			} else if c == "l" && position[0] > 0 {
				frontier = append(frontier, struct {
					path     string
					position [2]int
				}{working + "L", [2]int{position[0] - 1, position[1]}})
			} else if c == "r" && position[0] < 3 {
				frontier = append(frontier, struct {
					path     string
					position [2]int
				}{working + "R", [2]int{position[0] + 1, position[1]}})
			}
		}
	}

	fmt.Println()
	fmt.Println("Part2: What is the length of the longest path that reaches the vault?")
	fmt.Println(maxlen)
}
