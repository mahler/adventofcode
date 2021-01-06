package main

import "fmt"

func main() {
	cardKey := 16616892
	doorKey := 14505727

	cardKeyFound := false
	doorKeyFound := false
	doorKeyLoop := 0

	i := 0
	v := 1
	for {
		v *= 7
		v %= 20201227
		i++

		if v == cardKey {
			cardKeyFound = true
		}

		if v == doorKey {
			doorKeyLoop = i
			doorKeyFound = true
		}

		if cardKeyFound && doorKeyFound {
			break
		}
	}

	encryptionKey := 1
	for i := 0; i < doorKeyLoop; i++ {
		encryptionKey *= cardKey
		encryptionKey %= 20201227
	}

	fmt.Println()
	fmt.Println("Day 25")
	fmt.Println(encryptionKey)
}
