package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reports [][]int

	data, _ := os.Open("puzzle.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var report []int
		for _, field := range line {
			number, _ := strconv.Atoi(field)
			report = append(report, number)
		}
		reports = append(reports, report)
	}

	var safeReports int
	for _, report := range reports {
		if checkReport(report) {
			safeReports++
		}
	}
	fmt.Println("Part 1: How many reports are safe?")
	fmt.Println(safeReports)

	var safeReportsCompensated int
	for _, report := range reports {
		if checkReportCompensated(report) {
			safeReportsCompensated++
		}
	}
	fmt.Println()
	fmt.Println("Part 2:,  How many compensated reports are now safe?")
	fmt.Println(safeReportsCompensated)
}

func checkReport(report []int) bool {
	var increased int
	var decreased int
	for i := 0; i < (len(report) - 1); i++ {
		separatedSpaces := report[i : i+2]
		distance := separatedSpaces[0] - separatedSpaces[1]
		if distance == 0 {
			return false
		}
		if distance < 0 {
			decreased++
			distance = distance * -1
		} else {
			increased++
		}
		if distance > 3 {
			return false
		}
	}
	if (increased != 0) && (decreased != 0) {
		return false
	}
	return true
}

func checkReportCompensated(report []int) bool {
	if checkReport(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		reportCopy := make([]int, len(report))
		_ = copy(reportCopy, report)
		compensatedReport := append(reportCopy[:i], reportCopy[i+1:]...)

		if checkReport(compensatedReport) {
			return true
		}
	}
	return false
}
