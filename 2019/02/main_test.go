package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestRunProgram(t *testing.T) {
	testData := make(map[int]string)
	testData[3500] = "1,9,10,3,2,3,11,0,99,30,40,50"
	testData[2] = "1,0,0,0,99"
	//testData[6] = "2,3,0,3,99"
	//testData[9801] = "2,4,4,5,99,0"
	testData[30] = "1,1,1,4,99,5,6,0,99"

	for resultExpected, data := range testData {
		var intOpCode []int
		operations := strings.Split(string(data), ",")
		for _, value := range operations {
			opVal, _ := strconv.Atoi(value)
			intOpCode = append(intOpCode, opVal)
		}

		result, _ := runProgram(intOpCode)
		if result != resultExpected {
			t.Errorf("Wrong result for %v - got %v, wanted %v", data, result, resultExpected)
		}

	}

}
