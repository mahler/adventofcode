package main

import "fmt"

func main() {
	// To continue, please consult the code grid in the manual.  Enter the code at row 3010, column 3019.
	row := 3010
	col := 3019

	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 25: Let It Snow")

	num := 1
	for i := 0; i < col-1; i++ {
		num += i + 1
	}
	for j := 0; j < row; j++ {
		num += col - 1 + j
	}

	input := 20151125
	for k := 0; k < num-1; k++ {
		input = (input * 252533) % 33554393
	}

	fmt.Println("What code do you give the machine?")
	fmt.Println(input)
}
