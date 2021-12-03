package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	diagnosticMap := make(map[string]int)

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("DAY03, Part 1: Binary Diagnostic")
	fmt.Println("Total puzzle input: ", len(fileContent))

	for _, diagnostic := range fileContent {
		fmt.Println(diagnostic)
		for pos, binData := range diagnostic {
			fmt.Println(pos, "-", string(binData))
			mapKey := strconv.Itoa(pos) + "-" + string(binData)
			if _, ok := diagnosticMap[mapKey]; ok {
				diagnosticMap[mapKey]++
			} else {
				diagnosticMap[mapKey] = 1
			}
		}
	}

	//fmt.Println(diagnosticMap)
	maxPos := len(fileContent[0]) // Assume all rows are equal length
	gamma, epsilon := "", ""
	for x := 0; x < maxPos; x++ {
		bitOn := strconv.Itoa(x) + "-1"
		bitOff := strconv.Itoa(x) + "-0"

		if diagnosticMap[bitOn] > diagnosticMap[bitOff] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	fmt.Println("Bin Gamma", gamma)
	intGamma, _ := strconv.ParseInt(gamma, 2, 64)
	fmt.Println("Int Gamma", intGamma)

	fmt.Println("Bin Epsilon", epsilon)
	intEpsilon, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("Int Epsilon", intEpsilon)

	powerConsumption := intGamma * intEpsilon

	fmt.Println("power consumption:", powerConsumption)

	// ------------------------------------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")
}
