package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFIle("password.input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	strSlice := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Day 2, Part 1: Password Philosophy")
	totalPasswords := len(strSlice)
	passwordOk := 0
	passwordFail := 0

	// RegExp tester online: https://regoio.herokuapp.com/
	regexpData := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	for _, s1 := range strSlice {
		fields := regexpData.FindStringSubmatch(s1)
		// fmt.Println(fields)

		minCount, _ := strconv.Atoi(fields[1])
		maxCount, _ := strconv.Atoi(fields[2])
		letter := fields[3]
		password := fields[4]

		re := regexp.MustCompile(letter)
		matchSlice := re.FindAllString(password, -1)
		letterCount := len(matchSlice)

		if letterCount >= minCount && letterCount <= maxCount {
			//	fmt.Println("OK")
			passwordOk++
		} else {
			//	fmt.Println("NOT OK")
			passwordFail++
		}
		//		break
	}

	fmt.Println("Total passwords:", totalPasswords)
	fmt.Println("Password ok:", passwordOk)
	fmt.Println("Password fail:", passwordFail)
	fmt.Println()

	//
	//  PART TWO
	//

	pwdOk := 0
	pwdFail := 0

	fmt.Println("Day 2, Part 2: Positions in the password")
	for _, s2 := range strSlice {
		strs := strings.Split(s2, " ")
		letter := strs[1]
		password := strs[2]

		strs = strings.Split(strs[0], "-")
		pos1, _ := strconv.Atoi(strs[0])
		pos2, _ := strconv.Atoi(strs[1])
		letter = string(letter[0])

		checkLetter := 0
		if letter == string(password[pos1-1]) {
			checkLetter++
		}
		if letter == string(password[pos2-1]) {
			checkLetter++
		}

		if checkLetter == 1 {
			pwdOk++
		} else {
			pwdFail++
		}

	}

	fmt.Println("Total passwords:", totalPasswords)
	fmt.Println("Password ok:", pwdOk)
	fmt.Println("Password fail:", pwdFail)
}
