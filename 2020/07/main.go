package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func recurseColor(color string, colorMap map[string]map[string]int) []string {
	var returnColor []string

	for thisColor := range colorMap[color] {
		returnColor = append(returnColor, thisColor)

		subColor := recurseColor(thisColor, colorMap)
		returnColor = append(returnColor, subColor...)

	}
	return returnColor
}

func insideBagCount(color string, bagInBagMap map[string]map[string]int) int {
	totalBags := 1
	for thisColor, thisCount := range bagInBagMap[color] {
		// If the color in data was "no other bags", then the color field is "other" due to processing.
		if thisColor != "other" {
			insideBags := insideBagCount(thisColor, bagInBagMap)
			totalBags += (insideBags * thisCount)
			// fmt.Printf("'%v': %v * %v = %v\n", thisColor, insideBags, thisCount, (thisCount * insideBags))
		} else {
			totalBags = 1
		}

	}
	return totalBags
}

func main() {
	data, err := ioutil.ReadFile("baggage.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	records := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("DAY07, Part 1: Handy Haversacks")

	fmt.Println("Bagrules in dataset: ", len(records))

	bagColorContains := make(map[string]map[string]int)

	for _, bagRule := range records {
		bagColorIx := strings.Index(bagRule, " bags contain")
		bagColor := bagRule[0:bagColorIx]

		// fmt.Println("BagColor:", bagColor)

		// Add 14 to bagColorIx to skip pass " bags contain "
		bagContent := bagRule[bagColorIx+14:]
		// fmt.Println("*", bagContent)
		splitBagContent := strings.Split(bagContent, ",")
		for _, sbc := range splitBagContent {
			// remove preceeding or trailing spaces
			sbc = strings.Trim(sbc, ".")
			// remove number of bags and bag from string to just keep color
			containedColorBag := strings.TrimLeft(sbc, " 123456789")
			bagMark := strings.Index(containedColorBag, "bag")
			containedColorBag = containedColorBag[0:bagMark]
			ccBag := strings.Trim(containedColorBag, " ")

			if _, ok := bagColorContains[ccBag]; !ok {
				bagColorContains[ccBag] = make(map[string]int)

			}
			bagColorContains[ccBag][bagColor] = 1
		}
	}

	shinyGold := recurseColor("shiny gold", bagColorContains)
	sgMap := make(map[string]int)
	for _, uniqueColors := range shinyGold {
		sgMap[uniqueColors] = 1
	}
	fmt.Println("Bag colors which can contain a Shiny Gold Bag:", len(sgMap))

	//
	// PART 2
	//
	fmt.Println()
	fmt.Println("DAY07, Part 2: Bags in Bags")
	bagColorCount := make(map[string]map[string]int)

	for _, bagRule := range records {
		bagColorIx := strings.Index(bagRule, " bags contain")
		bagColor := bagRule[0:bagColorIx]
		//	fmt.Println("BagColor:", bagColor)
		bagContent := bagRule[bagColorIx+14 : len(bagRule)-1]
		//	fmt.Println("bC:", bagContent)
		splitContent := strings.Split(bagContent, ",")
		for _, spSco := range splitContent {
			// spSco has:  <number of bags> <bag color> bag(s)
			bagColorIx := strings.Index(spSco, "bag")
			spSco := spSco[0:bagColorIx]
			spSco = strings.Trim(spSco, " ")

			// spSco: "<number of bags> <name of color>"
			bagColorIx = strings.Index(spSco, " ")
			numBags, _ := strconv.Atoi(spSco[0:bagColorIx])
			colorBag := spSco[bagColorIx:]

			if _, ok := bagColorCount[bagColor]; !ok {
				bagColorCount[bagColor] = make(map[string]int)
			}
			colorBag = strings.Trim(colorBag, " ")
			bagColorCount[bagColor][colorBag] = numBags
		}

	}

	bagsCounted := insideBagCount("shiny gold", bagColorCount)
	// Don't count the shiny goldbag itself.
	bagsCounted--
	fmt.Println("Shiny Gold/ Bags in Bags total:", bagsCounted)
}
