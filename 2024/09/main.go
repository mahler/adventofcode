package main

import (
	"fmt"
	"os"
)

type Span struct {
	Index  int
	Sector int
	Size   int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)

	// P1
	files, free := parseInputP1(input)
	l, r := 0, len(files)-1

	for free[l].Sector < files[r].Sector {
		for i := 0; i < free[l].Size; i++ {
			files[r-i].Sector = free[l].Sector
		}
		r -= free[l].Size
		l++
	}

	fmt.Println("Part 1: Compact the amphipod's hard drive using the process he requested.")
	fmt.Println("What is the resulting filesystem checksum?")
	fmt.Println(checksum(files))

	// P2
	files, free = parseInputP2(input)

	for i := len(files) - 1; i >= 0; i-- {
		for j := range free {
			if free[j].Sector < files[i].Sector && free[j].Size >= files[i].Size {
				files[i].Sector = free[j].Sector
				free[j].Sector += files[i].Size
				free[j].Size -= files[i].Size
				break
			}
		}
	}

	fmt.Println()
	fmt.Println("P2: Start over, now compacting the amphipod's hard drive using this new method instead.")
	fmt.Println("What is the resulting filesystem checksum?")
	fmt.Println(checksum(files))
}

func parseInputP1(data string) ([]Span, []Span) {
	return parseInput(data, func(files *[]Span, index, pos, size int) {
		for k := 0; k < size; k++ {
			*files = append(*files, Span{
				Index:  index,
				Sector: pos + k,
				Size:   1,
			})
		}
	})
}

func parseInputP2(data string) ([]Span, []Span) {
	return parseInput(data, func(files *[]Span, index, pos, size int) {
		*files = append(*files, Span{
			Index:  index,
			Sector: pos,
			Size:   size,
		})
	})
}

func parseInput(data string, extend func(*[]Span, int, int, int)) ([]Span, []Span) {
	sector := 0
	var files, free []Span
	r := [](*[]Span){&files, &free}

	for i, c := range data {
		if v, ok := parse(c); ok {
			extend(r[i%2], i/2, sector, int(v))
			sector += int(v)
		}
	}

	return files, free
}

func parse(c rune) (uint8, bool) {
	if c >= '0' && c <= '9' {
		return uint8(c - '0'), true
	}
	return 0, false
}

func checksum(files []Span) int {
	total := 0
	for _, f := range files {
		for s := f.Sector; s < f.Sector+f.Size; s++ {
			total += f.Index * s
		}
	}
	return total
}
