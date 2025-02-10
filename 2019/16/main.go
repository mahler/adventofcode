package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readNumbers(input string) []int {
	input = strings.TrimSpace(input)
	result := make([]int, 0, len(input))
	for _, c := range input {
		num, _ := strconv.Atoi(string(c))
		result = append(result, num)
	}
	return result
}

func applyPhase(numbers []int, phases int) []int {
	pattern := []int{1, 0, -1, 0}
	next := make([]int, 0, len(numbers))

	for p := 0; p < phases; p++ {
		for idx1 := 0; idx1 < len(numbers); idx1++ {
			val1Result := 0
			for idx2 := idx1; idx2 < len(numbers); idx2++ {
				patternValue := pattern[((idx2-idx1)/(idx1+1))%4]
				val1Result += numbers[idx2] * patternValue
			}
			next = append(next, abs(val1Result)%10)
		}
		numbers, next = next, numbers
		next = next[:0]
	}
	return numbers
}

func foldToNumber(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result = result*10 + num
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input := "59712692690937920492680390886862131901538154314496197364022235676243731306353384700179627460533651346711155314756853419495734284609894966089975988246871687322567664499495407183657735571812115059436153203283165299263503632551949744441033411147947509168375383038493461562836199103303184064429083384309509676574941283043596285161244885454471652448757914444304449337194545948341288172476145567753415508006250059581738670546703862905469451368454757707996318377494042589908611965335468490525108524655606907405249860972187568380476703577532080056382150009356406585677577958020969940093556279280232948278128818920216728406595068868046480073694516140765535007"

	numbers := readNumbers(input)
	result := applyPhase(numbers, 100)
	fmt.Println("Part 1: After 100 phases of FFT, what are the first eight digits in the final output list?")
	fmt.Println(foldToNumber(result[:8]))

	// Part 2
	numbers = readNumbers(input)
	start := foldToNumber(numbers[:7])
	end := len(numbers) * 10000

	current := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		current = append(current, numbers[i%len(numbers)])
	}

	for p := 0; p < 100; p++ {
		sums := make([]int, len(current)+1)
		total := 0
		sums[0] = 0
		for i := 0; i < len(current); i++ {
			total += current[i]
			sums[i+1] = total
		}
		for i := 0; i < len(current); i++ {
			value := sums[len(sums)-1] - sums[i]
			current[i] = value % 10
		}
	}
	fmt.Println()
	fmt.Println("Part 2: After repeating your input signal 10000 times and running 100 phases of FFT,")
	fmt.Println("what is the eight-digit message embedded in the final output list?")
	fmt.Println(foldToNumber(current[:8]))
}
