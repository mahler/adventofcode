package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Adapted from https://github.com/nlowe/aoc2020/tree/master/challenge/day14
func main() {
	data, err := os.ReadFIle("program.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	dataInput := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println("Lines in dataInput:", len(dataInput))
	fmt.Println()

	fmt.Println("Day 14, Part 1: Docking Data")
	var andMask int64 = 0
	var orMask int64 = 0

	mem := map[int]int64{}

	for _, op := range dataInput {
		parts := strings.Split(op, " = ")

		if parts[0] == "mask" {
			andMask, _ = strconv.ParseInt(strings.ReplaceAll(parts[1], "X", "1"), 2, 0)
			orMask, _ = strconv.ParseInt(strings.ReplaceAll(parts[1], "X", "0"), 2, 0)
		} else {
			addr := mustAtoi(strings.TrimSuffix(strings.TrimPrefix(parts[0], "mem["), "]"))
			value := int64(mustAtoi(parts[1]))

			value &= andMask
			value |= orMask

			mem[addr] = value
		}
	}

	memorySum := 0
	for _, v := range mem {
		memorySum += int(v)
	}

	fmt.Println("The memory sum upon completion is:", memorySum)

	fmt.Println()
	fmt.Println("Part 2")

	mask := ""
	mem2 := map[string]int{}

	for _, op := range dataInput {
		parts := strings.Split(op, " = ")

		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			addr := mustAtoi(strings.TrimSuffix(strings.TrimPrefix(parts[0], "mem["), "]"))
			value := mustAtoi(parts[1])

			for _, generatedAddr := range permuteMask("", mask, fmt.Sprintf("%036b", addr)) {
				mem2[generatedAddr] = value
			}
		}
	}

	memoryAddressDecoder := 0
	for _, v := range mem2 {
		memoryAddressDecoder += v
	}

	fmt.Println("Memory address decoding value:", memoryAddressDecoder)
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func permuteMask(mask, remaining, addr string) []string {
	if len(remaining) == 0 {
		return []string{mask}
	}

	switch remaining[0] {
	case '0':
		return permuteMask(mask+string(addr[len(mask)]), remaining[1:], addr)
	case '1':
		return permuteMask(mask+"1", remaining[1:], addr)
	case 'X':
		return append(permuteMask(mask+"0", remaining[1:], addr), permuteMask(mask+"1", remaining[1:], addr)...)
	default:
		panic(fmt.Errorf("unknown bitmask type %s", string(remaining[0])))
	}
}
