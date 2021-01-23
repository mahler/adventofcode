package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {

	input := "iwrupvqb"
	counter := 0

	fmt.Println()
	fmt.Print("Day 04, Part 1: The Ideal Stocking Stuffer")
	for {
		data := input + strconv.Itoa(counter)
		sum := md5.Sum([]byte(data))
		str := hex.EncodeToString(sum[:])
		if str[0:5] == "00000" {
			// AdventCoin found.
			fmt.Println(counter)
			break
		}
		counter++
	}
	fmt.Println("lowest positive number - zero5:", counter)
	// ------------ PART 2 ------------------------

	fmt.Println()
	fmt.Print("Day 04, Part 2:")
	// reset counter for good measure.
	counter = 0
	for {
		data := input + strconv.Itoa(counter)
		sum := md5.Sum([]byte(data))
		str := hex.EncodeToString(sum[:])
		if str[0:6] == "000000" {
			// AdventCoin found.
			fmt.Println(counter)
			break
		}
		counter++
	}
	fmt.Println("lowest positive number - zero6:", counter)
}
