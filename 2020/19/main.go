package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	data, err := os.ReadFIle("puzzle.sample")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	rules := make(map[string]string)
	var messsages []string

	ruleSection := true
	for _, fileLine := range fileLines {
		if fileLine == "" {
			ruleSection = false
		} else if ruleSection {
			ixNumber := strings.Index(fileLine, ":")
			ruleNumber := fileLine[:ixNumber]
			ruleContent := strings.TrimSpace(fileLine[ixNumber+1:])

			rules[ruleNumber] = ruleContent
		} else {
			messsages = append(messsages, fileLine)
		}
	}

	fmt.Println("Validation rules:", len(rules))
	for k, v := range rules {
		fmt.Println(k, ":", v)
	}
	fmt.Println("Data messages:", len(messsages))

	fmt.Println()
	fmt.Println("Day 19, Part 1: Monster Messages")
	messagesMatch := 0
	for _, msg := range messsages {
		if checkRule(msg, "0", rules) {
			fmt.Println("Checked", msg, "match!")
			messagesMatch++
		}
		break
	}

	fmt.Println("messages completely match:", messagesMatch)
}

func checkRule(input string, ruleID string, rules map[string]string) bool {
	thisRule := rules[ruleID]
	fmt.Println("checkRule/ data:", input, " // rule:", ruleID, "(", thisRule, ")")
	if strings.Contains(thisRule, "|") {
		fmt.Println("subrules in play")
		subRules := strings.Split(thisRule, "|")
		for _, subRule := range subRules {
			singleRules := strings.Split(subRule, " ")

			fmt.Print("--", subRule, "//", singleRules)

			return true
		}
	} else if strings.Contains(thisRule, "\"") {
		// rule is single letter
		letter := string(thisRule[1])
		if strings.Contains(input, letter) {
			return true
		}
	} else {
		// rule is series of numbers.
		split := strings.Split(thisRule, " ")

		for _, subRule := range split {
			checkOk := checkRule(input, subRule, rules)
			if !checkOk {
				return false
			}
		}
		return true

	}
	return false
}
