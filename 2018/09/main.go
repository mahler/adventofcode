package main

import (
	"fmt"
)

var metadataTotal int

func main() {
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("What is the winning Elf's score?")
	highScore := getHighestScore(473, 70904)
	fmt.Println(highScore)

	//-----
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("What would the new winning Elf's score be")
	fmt.Println("if the number of the last marble were 100 times larger?")
	highScore = getHighestScore(473, 70904*100)
	fmt.Println(highScore)
}

// Marble creates linked list that makes a circle
type Marble struct {
	Left  *Marble
	Right *Marble
	Value int
}

// Source: https://github.com/thlacroix/goadvent/blob/master/2018/day09/main.go
func getHighestScore(playerCount, lastMarble int) int {
	// creating the initial marble
	initialMarble := &Marble{Value: 0}
	initialMarble.Left = initialMarble
	initialMarble.Right = initialMarble
	currentMarble := initialMarble
	// keeping the score for each player
	scores := make(map[int]int)
	for i := 1; i <= lastMarble; i++ {
		currentPlayer := (i-1)%playerCount + 1

		// However, if the marble that is about to be placed has a number which is a multiple of 23,
		// something entirely different happens. First, the current player keeps the marble they would
		// have placed, adding it to their score. In addition, the marble 7 marbles counter-clockwise
		// from the current marble is removed from the circle and also added to the current player's score.
		// The marble located immediately clockwise of the marble that was removed becomes the new current marble.
		if i != 0 && i%23 == 0 {
			for j := 0; j < 7; j++ {
				currentMarble = currentMarble.Left
			}
			currentMarble.Left.Right = currentMarble.Right
			currentMarble.Right.Left = currentMarble.Left
			scores[currentPlayer] += i + currentMarble.Value
			currentMarble = currentMarble.Right
		} else {
			// Then, each Elf takes a turn placing the lowest-numbered remaining marble into the circle between the
			// marbles that are 1 and 2 marbles clockwise of the current marble. (When the circle is large enough,
			// this means that there is one marble between the marble that was just placed and the current marble.)
			// The marble that was just placed then becomes the current marble.
			beforeNewMarble := currentMarble.Right
			afterNewMarble := beforeNewMarble.Right
			currentMarble = &Marble{Value: i, Left: beforeNewMarble, Right: afterNewMarble}
			beforeNewMarble.Right = currentMarble
			afterNewMarble.Left = currentMarble
		}
	}
	// returning max score
	var max int
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}
