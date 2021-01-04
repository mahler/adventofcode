package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	// Read instructions
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	fileLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	// Read souce file and setup ingredientes
	ingredients := map[int]ingredient{}
	rxIngredients := regexp.MustCompile("^([a-zA-Z]+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)$")

	for i, line := range fileLines {
		fields := rxIngredients.FindStringSubmatch(line)
		name := fields[1]
		capacity, _ := strconv.Atoi(fields[2])
		durability, _ := strconv.Atoi(fields[3])
		flavor, _ := strconv.Atoi(fields[4])
		texture, _ := strconv.Atoi(fields[5])
		calories, _ := strconv.Atoi(fields[6])
		ingredients[i] = ingredient{name, capacity, durability, flavor, texture, calories}
	}

	// Part 1
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 15: Science for Hungry People")

	partitions := [][]int{}
	part([]int{}, 100, len(ingredients), &partitions)
	first := true
	highestScore := 0
	second := true
	highestCalorie := 0
	for i := 0; i < len(partitions); i++ {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0 // p2 only

		for p := 0; p < len(ingredients); p++ {
			capacity += partitions[i][p] * ingredients[p].capacity
			durability += partitions[i][p] * ingredients[p].durability
			flavor += partitions[i][p] * ingredients[p].flavor
			texture += partitions[i][p] * ingredients[p].texture
			calories += partitions[i][p] * ingredients[p].calories // p2 only

		}
		// Part 1 calc
		cookieScore := 0
		if capacity > 0 && durability > 0 && flavor > 0 && texture > 0 {
			cookieScore = capacity * durability * flavor * texture
		}
		if first || cookieScore > highestScore {
			highestScore = cookieScore
			first = false
		}
		// Part 2 extension...
		calorieScore := 0
		if capacity > 0 && durability > 0 && flavor > 0 && texture > 0 && calories == 500 {
			calorieScore = capacity * durability * flavor * texture
		}
		if second || calorieScore > highestCalorie {
			highestCalorie = calorieScore
			second = false
		}

	}
	fmt.Println("what is the total score of the highest-scoring cookie?")
	fmt.Println(highestScore)
	// -------------------------
	fmt.Println("Part 2: what is the total score of the highest-scoring")
	fmt.Println("cookie you can make with a calorie total of 500?")
	fmt.Println(highestCalorie)

}

func part(prefix []int, number int, parts int, results *[][]int) {
	if parts == 1 {
		newPrefix := append(append([]int{}, prefix...), number)
		*results = append(*results, [][]int{newPrefix}...)
	} else {
		for i := 0; i <= number; i++ {
			newPrefix := append(append([]int{}, prefix...), i)
			part(newPrefix, number-i, parts-1, results)
		}
	}
}
