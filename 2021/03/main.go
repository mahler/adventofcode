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

	for _, diagnostic := range fileContent {
		//fmt.Println(diagnostic)
		for pos, binData := range diagnostic {
			// fmt.Println(pos, "-", string(binData))
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

	fmt.Println("Power consumption:", powerConsumption)

	// ------------------------------------------------------------------------
	fmt.Println()
	fmt.Println("Part 2")

	oxygenRows := make([]string, len(fileContent))
	copy(oxygenRows, fileContent)

	for pos := 0; pos < maxPos; pos++ {
		// Count ones and zeros
		ones, zeros := 0, 0
		for _, val := range oxygenRows {
			if val[pos] == '1' {
				ones++
			} else if val[pos] == '0' {
				zeros++
			}
		}

		bitSet := '0'
		// If 0 and 1 are equally common, keep values with a 1 in the position being considered.
		if ones >= zeros {
			bitSet = '1'
		}

		for key, val := range oxygenRows {
			if val[pos] != byte(bitSet) {
				oxygenRows[key] = "--------------------------------"
			}
		}
	}
	oxygenGeneratorRating := ""
	for _, val := range oxygenRows {
		if val[0] != '-' {
			oxygenGeneratorRating = val
			break
		}
	}

	fmt.Println("oxygen generator rating (bin):", oxygenGeneratorRating)
	intoxygenGeneratorRating, _ := strconv.ParseInt(string(oxygenGeneratorRating), 2, 64)
	fmt.Println("oxygen generator rating (dec):", intoxygenGeneratorRating)
	fmt.Println()

	// Moving on to CO2 scrubber..
	co2Rows := make([]string, len(fileContent))
	copy(co2Rows, fileContent)

	zapped := 0
	for pos := 0; pos < maxPos; pos++ {
		// Count ones and zeros
		ones, zeros := 0, 0
		for _, val := range co2Rows {
			if val[pos] == '1' {
				ones++
			} else if val[pos] == '0' {
				zeros++
			}
		}

		bitSet := '0'
		// If 0 and 1 are equally common, keep values with a 0 in the position being considered.
		if ones < zeros {
			bitSet = '1'
		}

		for key, val := range co2Rows {
			if val[pos] != '-' && val[pos] != byte(bitSet) {
				co2Rows[key] = "--------------------------------"
				zapped++
			}
		}
		// Break if only one item left.
		if zapped > len(co2Rows)-2 {
			break
		}
	}

	co2scrubbing := ""
	for _, val := range co2Rows {
		if val[0] != '-' {
			co2scrubbing = val
			break
		}
	}
	fmt.Println("CO2 scrubbing (bin):", co2scrubbing)
	intco2scrubbing, _ := strconv.ParseInt(co2scrubbing, 2, 64)
	fmt.Println("CO2 scrubbing (dec):", intco2scrubbing)
	fmt.Println()

	fmt.Println("What is the life support rating of the submarine?")
	fmt.Println(intoxygenGeneratorRating * intco2scrubbing)
}
