package main

import (
	"fmt"
	"os"
	"strings"
)

func count[T comparable](v []T, x T) int {
	count := 0
	for _, n := range v {
		if n == x {
			count++
		}
	}
	return count
}

func parse(input string) []uint8 {
	pixels := make([]uint8, 0, len(input))
	for _, ch := range input {
		if digit := uint8(ch - '0'); digit >= 0 && digit <= 9 {
			pixels = append(pixels, digit)
		}
	}
	return pixels
}

func merge(layers [][]uint8) []uint8 {
	result := make([]uint8, 0, len(layers[0]))
	iterators := make([]int, len(layers))

	for {
		var pixel *uint8
		for i, layer := range layers {
			if iterators[i] >= len(layer) {
				return result
			}

			digit := layer[iterators[i]]
			iterators[i]++

			if digit != 2 && pixel == nil {
				pixel = &digit
			}
		}

		if pixel == nil {
			break
		}
		result = append(result, *pixel)
	}

	return result
}

func main() {
	width := 25
	height := 6

	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Failed to read input file:", err)
		return
	}

	input := strings.TrimSpace(string(inputBytes))
	pixels := parse(input)

	var layers [][]uint8
	for i := 0; i < len(pixels); i += width * height {
		end := i + width*height
		if end > len(pixels) {
			end = len(pixels)
		}
		layers = append(layers, pixels[i:end])
	}

	// Part 1
	var leastZerosLayer []uint8
	minZeros := len(layers[0]) + 1

	for _, layer := range layers {
		zeros := count(layer, uint8(0))
		if zeros < minZeros {
			minZeros = zeros
			leastZerosLayer = layer
		}
	}

	fmt.Println("Part 1: On that layer, what is the number of 1 digits multiplied by the number of 2 digits?")
	fmt.Println(count(leastZerosLayer, uint8(1)) * count(leastZerosLayer, uint8(2)))

	// Part 2
	merged := merge(layers)
	var lines []string

	for i := 0; i < len(merged); i += 25 {
		var line strings.Builder
		for j := 0; j < 25 && i+j < len(merged); j++ {
			if merged[i+j] == 0 {
				line.WriteString("  ")
			} else {
				line.WriteString("##")
			}
		}
		lines = append(lines, line.String())
	}

	fmt.Println()
	fmt.Println("Part 2: What message is produced after decoding your image?")
	fmt.Println()
	fmt.Println(strings.Join(lines, "\n"))

}
