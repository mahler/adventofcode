package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFIle("puzzle.input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//   fmt.Println("Contents of file:", string(data))

	strSlice := strings.Split(strings.TrimSpace(string(data)), "\n")
	//   fmt.Println(strSlice)

	fmt.Println("PART 1 - two numbers which add up to 2020")
	for _, s1 := range strSlice {
		for _, s2 := range strSlice {
			i1, _ := strconv.Atoi(s1)
			i2, _ := strconv.Atoi(s2)
			//fmt.Printf("%v +  %v = %v \n", s1, s2, (i1 + i2))
			//fmt.Println(i1 + i2)
			var sum int
			if sum = i1 + i2; sum == 2020 {
				fmt.Printf("Sum: %v +  %v = %v \n", s1, s2, (i1 + i2))
				fmt.Printf("Multiply: %v\n\n", i1*i2)
				goto Exit1
			}
		}
	}
Exit1:

	// Part 2
	fmt.Println("PART 2 - three numbers which add up to 2020")
	for _, s1 := range strSlice {
		for _, s2 := range strSlice {
			for _, s3 := range strSlice {
				i1, _ := strconv.Atoi(s1)
				i2, _ := strconv.Atoi(s2)
				i3, _ := strconv.Atoi(s3)

				// fmt.Printf("%v +  %v + %v = %v \n", i1, i2, i3, (i1 + i2 + i3))
				//fmt.Println(i1 + i2 + i3)
				var sum int
				if sum = i1 + i2 + i3; sum == 2020 {
					fmt.Printf("%v +  %v + %v = %v \n", i1, i2, i3, (i1 + i2 + i3))
					fmt.Printf("Multiply: %v\n", i1*i2*i3)
					goto Exit2
				}

			}
		}
	}
Exit2:
}
