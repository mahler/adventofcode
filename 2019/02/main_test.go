package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestRunProgram(t *testing.T) {
	m2fTestData := make(map[int]string)
	m2fTestData[3500] = "1,9,10,3,2,3,11,0,99,30,40,50"
	m2fTestData[2] = "1,0,0,0,99"
	m2fTestData[6] = "2,3,0,3,99"
	m2fTestData[9801] = "2,4,4,5,99,0"
	m2fTestData[30] = "1,1,1,4,99,5,6,0,99"

	for resultExpected, data := range m2fTestData {
		fmt.Println("*")
		var intOpCode []int
		operations := strings.Split(string(data), ",")
		for _, value := range operations {
			opVal, _ := strconv.Atoi(value)
			intOpCode = append(intOpCode, opVal)
		}

		result := RunProgram(intOpCode)
		if result != resultExpected {
			t.Errorf("Wrong result for %v - got %v, wanted %v", data, result, resultExpected)
		}

	}

}
