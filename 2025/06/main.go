package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func product(nums []int) int {
	p := 1
	for _, n := range nums {
		p *= n
	}
	return p
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func hasNumeric(s string) bool {
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			return true
		}
	}
	return false
}

func part1(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var rows [][]interface{}
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		if hasNumeric(line) {
			fields := strings.Fields(line)
			var row []interface{}
			for _, field := range fields {
				if n, err := strconv.Atoi(field); err == nil {
					row = append(row, n)
				} else {
					row = append(row, field)
				}
			}
			rows = append(rows, row)
		} else {
			fields := strings.Fields(line)
			var row []interface{}
			for _, field := range fields {
				row = append(row, field)
			}
			rows = append(rows, row)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Transpose
	if len(rows) == 0 {
		fmt.Println("Part 1: 0")
		return nil
	}
	
	problems := make([][]interface{}, len(rows[0]))
	for j := range problems {
		problems[j] = make([]interface{}, len(rows))
		for i := range rows {
			problems[j][i] = rows[i][j]
		}
	}

	checksum := 0
	for _, problem := range problems {
		lastIdx := len(problem) - 1
		if op, ok := problem[lastIdx].(string); ok && op == "*" {
			var nums []int
			for i := 0; i < lastIdx; i++ {
				if n, ok := problem[i].(int); ok {
					nums = append(nums, n)
				}
			}
			checksum += product(nums)
		} else {
			var nums []int
			for i := 0; i < lastIdx; i++ {
				if n, ok := problem[i].(int); ok {
					nums = append(nums, n)
				}
			}
			checksum += sum(nums)
		}
	}

	fmt.Printf("Part 1: %d\n", checksum)
	return nil
}

func part2(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var rows []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasSuffix(line, "\n") {
			line += " "
		}
		rows = append(rows, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if len(rows) == 0 {
		fmt.Println("Part 2: 0")
		return nil
	}

	var humanMathList [][]string
	var numList []string

	for j := len(rows[0]) - 1; j >= 0; j-- {
		allSpace := true
		for r := 0; r < len(rows); r++ {
			if j < len(rows[r]) && !unicode.IsSpace(rune(rows[r][j])) {
				allSpace = false
				break
			}
		}

		if allSpace {
			if len(numList) > 0 {
				humanMathList = append(humanMathList, numList)
				numList = nil
			}
			continue
		}

		var num string
		for i := 0; i < len(rows); i++ {
			if j >= len(rows[i]) {
				continue
			}
			ch := rows[i][j]
			if unicode.IsDigit(rune(ch)) {
				num += string(ch)
			}
			if ch == '+' || ch == '*' {
				numList = append([]string{string(ch)}, numList...)
			}
		}
		
		numList = append(numList, num)
		
		if j == 0 && len(numList) > 0 {
			humanMathList = append(humanMathList, numList)
		}
	}

	checksum := 0
	for _, problem := range humanMathList {
		if len(problem) == 0 {
			continue
		}
		
		if problem[0] == "+" {
			var nums []int
			for i := 1; i < len(problem); i++ {
				if n, err := strconv.Atoi(problem[i]); err == nil {
					nums = append(nums, n)
				}
			}
			checksum += sum(nums)
		} else {
			var nums []int
			for i := 1; i < len(problem); i++ {
				if n, err := strconv.Atoi(problem[i]); err == nil {
					nums = append(nums, n)
				}
			}
			checksum += product(nums)
		}
	}

	fmt.Printf("Part 2: %d\n", checksum)
	return nil
}

func main() {
	filename := "input.txt"
	
	if err := part1(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Part 1 error: %v\n", err)
	}
	
	if err := part2(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Part 2 error: %v\n", err)
	}
}
