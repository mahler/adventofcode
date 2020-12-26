package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getSeatID(strSeat string) int {

	rowNumber := 0
	columnNumber := 0

	// Find row number
	if strSeat[0:1] == "B" {
		rowNumber += 64
	}
	if strSeat[1:2] == "B" {
		rowNumber += 32
	}
	if strSeat[2:3] == "B" {
		rowNumber += 16
	}
	if strSeat[3:4] == "B" {
		rowNumber += 8
	}
	if strSeat[4:5] == "B" {
		rowNumber += 4
	}
	if strSeat[5:6] == "B" {
		rowNumber += 2
	}
	if strSeat[6:7] == "B" {
		rowNumber++
	}

	// Find column number
	if strSeat[7:8] == "R" {
		columnNumber += 4
	}
	if strSeat[8:9] == "R" {
		columnNumber += 2
	}
	if strSeat[9:10] == "R" {
		columnNumber++
	}
	return rowNumber*8 + columnNumber
}

func main() {
	data, err := ioutil.ReadFile("boarding.pass")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//   fmt.Println("Contents of file:", string(data))

	strSlice := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Records read:", len(strSlice))

	fmt.Println()
	fmt.Println("Day 05, Part 1:  Day 5: Binary Boarding")

	maxSeatID := 0

	var seatMap = make(map[int]string)

	for _, seatStr := range strSlice {
		seatID := getSeatID(seatStr)

		// Save for use in part 2
		seatMap[seatID] = seatStr

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println("Highest SeatID found is", maxSeatID)

	fmt.Println()
	fmt.Println("Day 05, Part 2: Finding seat")

	mySeat := 0

	// Max in loop set by result from part 1
	for i := 1; i < 866; i++ {
		beforeNumber := i - 1
		afterNumber := i + 1

		checkFailed := 0
		if _, ok := seatMap[beforeNumber]; !ok {
			checkFailed++
		}
		if _, ok := seatMap[afterNumber]; !ok {
			checkFailed++
		}
		if _, ok := seatMap[i]; ok {
			checkFailed++
		} else {
			// If checkFailed is zero, then there is a seatID before and after.
			if checkFailed == 0 {
				mySeat = i
			}
		}
	}
	fmt.Println("My seatID is", mySeat)

	fmt.Println()
}
