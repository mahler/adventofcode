package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	children []*node
	metadata []int
}

func main() {
	fileContent, _ := os.ReadFIle("puzzle.txt")
	split := strings.Split(string(fileContent), " ")
	numbers := make([]int, len(split))
	for i, s := range split {
		numbers[i], _ = strconv.Atoi(s)
	}
	// Recursively unpack numbers...
	root, _ := parse(numbers)

	// PArt 1 ----------------------------
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("Day 08, Part 1: Memory Maneuver")
	metadataSum := sum(root)
	fmt.Println("What is the sum of all metadata entries?")
	fmt.Println(metadataSum)

	// Part 2 --------------------------
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("What is the value of the root node?")
	fmt.Println(root.value())
}

// parse the input slice into tree
func parse(input []int) (*node, int) {
	addToIndex := 2
	n := &node{}
	for k := input[0]; k > 0; k-- {
		child, newAdd := parse(input[addToIndex:])
		addToIndex += newAdd
		n.children = append(n.children, child)
	}
	n.metadata = input[addToIndex : addToIndex+input[1]]
	addToIndex += input[1]
	return n, addToIndex
}

// Sum findes the sum of all metadata entries
func sum(n *node) int {
	su := 0
	for _, c := range n.children {
		su += sum(c)
	}
	for _, meta := range n.metadata {
		su += meta
	}
	return su
}

// value calculate the value recursicely of a node
func (n *node) value() int {
	sum := 0
	if len(n.children) == 0 {
		for _, mData := range n.metadata {
			sum += mData
		}
		return sum
	}
	for _, mData := range n.metadata {
		if mData > len(n.children) {
			sum += 0
		} else {
			sum += n.children[mData-1].value()
		}
	}
	return sum
}
