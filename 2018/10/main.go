package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Starlight struct {
	posX      int
	posY      int
	velocityX int
	velocityY int
}

func main() {

	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")
	fmt.Println()
	fmt.Println("2018")
	fmt.Println("DAY10, Part 1: The Stars Align")

	stars := []*Starlight{}
	rxStar := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)

	for _, starData := range fileContent {
		result := rxStar.FindStringSubmatch(starData)
		newStar := Starlight{}
		newStar.posX, _ = strconv.Atoi(result[1])
		newStar.posY, _ = strconv.Atoi(result[2])
		newStar.velocityX, _ = strconv.Atoi(result[3])
		newStar.velocityY, _ = strconv.Atoi(result[4])

		stars = append(stars, &newStar)
	}

	round := 1
	for {
		for _, star := range stars {
			star.posX += star.velocityX
			star.posY += star.velocityY
		}

		_, minY, _, maxY := edges(stars)
		//	fmt.Println(round, ": Cols/", maxY-minY)

		if maxY-minY < 10 {
			fmt.Println()
			printCanvas(stars)

			fmt.Println()
			fmt.Println("Part2:")
			fmt.Println("Round", round)
		} else {
			//fmt.Println("Round", round)
		}

		round++
		if round > 10005 {
			break
		}
	}
}

func printCanvas(s []*Starlight) {
	starmap := make(map[int]map[int]bool)

	minX, minY, maxX, maxY := edges(s)

	for row := minY; row <= maxY; row++ {
		starmap[row] = make(map[int]bool)
		for col := minX; col <= maxX; col++ {
			starmap[row][col] = false

		}
	}

	for _, star := range s {
		starmap[star.posY][star.posX] = true
	}

	for row := minY; row <= maxY; row++ {
		fmt.Printf("%3d : ", row)
		for col := minX; col <= maxX; col++ {
			if starmap[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func edges(vectors []*Starlight) (int, int, int, int) {
	minX, minY, maxX, maxY := 1000, 1000, -1000, -1000
	for _, v := range vectors {
		if v.posX > maxX {
			maxX = v.posX
		} else if v.posX < minX {
			minX = v.posX
		}
		if v.posY > maxY {
			maxY = v.posY
		} else if v.posY < minY {
			minY = v.posY
		}
	}
	return minX, minY, maxX, maxY
}
