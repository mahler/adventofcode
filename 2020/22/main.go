package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type crabCards struct {
	cards []int
}

func main() {
	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	strSlice := strings.Split(strings.TrimSpace(string(data)), "\n")

	playerNumber := 1

	player1 := []int{}
	player2 := []int{}

	for _, row := range strSlice {
		if len(row) == 0 {
			playerNumber++
		} else if string(row[0]) == "P" {
			// player(1|2)
			continue
		} else {
			cardNumber, _ := strconv.Atoi(row)

			if playerNumber == 1 {
				player1 = append(player1, cardNumber)
			} else {
				player2 = append(player2, cardNumber)
			}
		}
	}

	fmt.Println()
	fmt.Println("2020")
	fmt.Println("Day 22, Part 1: Crab Combat")
	round := 1
	for {

		// p1 draw card
		p1card := player1[0]
		player1 = player1[1:]

		// p2 draw card
		p2card := player2[0]
		player2 = player2[1:]

		//		fmt.Println("Round", round, "/ P1 picked:", p1card, "* P2 picked:", p2card)
		if p1card > p2card {
			player1 = append(player1, p1card, p2card)
		} else if p2card > p1card {
			player2 = append(player2, p2card, p1card)
		}

		if len(player1) == 0 || len(player2) == 0 {
			// end game
			break
		}
		round++
	}

	// Find winner & calc score
	if len(player1) > 0 {
		fmt.Println("Player 1 won with a score for", calcScore(player1))
	} else {
		fmt.Println("Player 2 won with a score for", calcScore(player2))
	}
}

func calcScore(cards []int) int {
	score := 0
	for i, val := range cards {
		score += (len(cards) - i) * val
	}
	return score
}
