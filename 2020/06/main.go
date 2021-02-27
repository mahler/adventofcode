package main

import (
	"fmt"
	"strings"
)

func main() {
	data, err := os.ReadFIle("customs.data")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	records := strings.Split(string(data), "\n\n")
	fmt.Println()
	fmt.Println("DAY06, Part 1: Custom Customs")
	fmt.Println("Total groups in dataset: ", len(records))

	var groupYes = make(map[int]int)

	for i, groupData := range records {
		var groupAnswers = make(map[string]int)

		groupRecord := strings.Split(groupData, "\n")
		for _, data := range groupRecord {
			for k := 0; k < len(data); k++ {
				groupAnswers[data[k:k+1]] = 1
			}
		}
		//fmt.Println("Group", i, " Answers:", len(groupAnswers))
		groupYes[i] = len(groupAnswers)
	}

	yesGroupSum := 0
	for i := 0; i < len(groupYes); i++ {
		yesGroupSum += groupYes[i]
	}

	fmt.Println("Yes sum across groups:", yesGroupSum)
	fmt.Println()

	//
	// PART 2
	//
	fmt.Println("DAY06, Part 2: All answer yes")

	totalAnswersPresent := 0
	for _, groupData := range records {

		var groupAnswers = make(map[string]int)
		groupRecord := strings.Split(groupData, "\n")

		for _, answersFromSingle := range groupRecord {
			// Find unique answers and build a map of it submitted answers in group
			for k := 0; k < len(answersFromSingle); k++ {

				thisLetter := answersFromSingle[k : k+1]
				// Keep this to know the unique answers for all in group.
				groupAnswers[thisLetter] = 1
			}
		}

		for checkletter := range groupAnswers {
			letterFoundInAnswer := 0
			//fmt.Println("checking", checkletter)
			for _, answersFromSingle := range groupRecord {
				// fmt.Println(answersFromSingle)
				if strings.Index(answersFromSingle, checkletter) >= 0 {
					letterFoundInAnswer++
				}
			}
			//			fmt.Println("Found", checkletter, "in", letterFoundInAnswer, "out of", len(groupRecord))
			if letterFoundInAnswer == len(groupRecord) {
				totalAnswersPresent++
			}

		}

	}
	fmt.Println("Total Answers Present from all in group:", totalAnswersPresent)
}
