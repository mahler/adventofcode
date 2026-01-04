package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func validIDs(ids []int, ranges []Range) int {
	i := 0
	j := 0
	freshIDs := 0

	for j < len(ids) && i < len(ranges) {
		if ids[j] > ranges[i].End {
			i++
		} else if ids[j] < ranges[i].Start {
			j++
		} else {
			freshIDs++
			j++
		}
	}

	return freshIDs
}

func dedup(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	fixed := []Range{ranges[0]}
	prev := &fixed[0]

	for _, r := range ranges[1:] {
		if r.End <= prev.End {
			continue
		}
		if r.Start <= prev.End {
			r.Start = prev.End + 1
		}
		fixed = append(fixed, r)
		prev = &fixed[len(fixed)-1]
	}

	return fixed
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	if len(parts) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid input format\n")
		os.Exit(1)
	}

	freshInput := strings.Fields(parts[0])
	idsInput := strings.Fields(parts[1])

	// Parse ranges
	ranges := make([]Range, 0, len(freshInput))
	for _, row := range freshInput {
		rangeParts := strings.Split(row, "-")
		if len(rangeParts) != 2 {
			continue
		}
		start, err1 := strconv.Atoi(rangeParts[0])
		end, err2 := strconv.Atoi(rangeParts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	// Parse IDs
	ids := make([]int, 0, len(idsInput))
	for _, idStr := range idsInput {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	// Sort both slices
	sort.Ints(ids)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	ranges = dedup(ranges)

	fmt.Println("How many of the available ingredient IDs are fresh?")
	fmt.Println(validIDs(ids, ranges))

	total := 0
	for _, r := range ranges {
		total += r.End - r.Start + 1
	}

	fmt.Println()
	fmt.Println("How many ingredient IDs are considered to be fresh according to the fresh ingredient ID ranges?")
	fmt.Println(total)
}
