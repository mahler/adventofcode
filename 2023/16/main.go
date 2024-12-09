package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	DIRS    = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	DNS     = []string{"R", "L", "D", "U"}
	MIRRORS = map[rune]map[string][]string{
		'.':  {"R": {"R"}, "L": {"L"}, "D": {"D"}, "U": {"U"}},
		'-':  {"R": {"R"}, "L": {"L"}, "D": {"L", "R"}, "U": {"L", "R"}},
		'|':  {"R": {"D", "U"}, "L": {"D", "U"}, "D": {"D"}, "U": {"U"}},
		'/':  {"R": {"U"}, "L": {"D"}, "D": {"L"}, "U": {"R"}},
		'\\': {"R": {"D"}, "L": {"U"}, "D": {"R"}, "U": {"L"}},
	}
)

func countFrom(ll [][]rune, start []interface{}) int {
	illum := make(map[[3]interface{}]bool)
	var illuminate func(int, int, string)
	illuminate = func(x, y int, dr string) {
		key := [3]interface{}{x, y, dr}
		if illum[key] {
			return
		}
		illum[key] = true
		mr := ll[x][y]
		for _, nxt := range MIRRORS[mr][dr] {
			dirIndex := indexOf(DNS, nxt)
			nxtDir := DIRS[dirIndex]
			nx := x + nxtDir[0]
			ny := y + nxtDir[1]
			if nx >= 0 && nx < len(ll) && ny >= 0 && ny < len(ll[0]) {
				illuminate(nx, ny, nxt)
			}
		}
	}
	illuminate(start[0].(int), start[1].(int), start[2].(string))

	unique := make(map[[2]int]bool)
	for k := range illum {
		unique[[2]int{k[0].(int), k[1].(int)}] = true
	}
	return len(unique)
}

func indexOf(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	ll := make([][]rune, len(lines))
	for i, line := range lines {
		ll[i] = []rune(line)
	}

	tests := [][]interface{}{}
	for x := 0; x < len(ll); x++ {
		tests = append(tests, []interface{}{x, 0, "R"})
		tests = append(tests, []interface{}{x, len(ll[0]) - 1, "L"})
	}
	for y := 0; y < len(ll[0]); y++ {
		tests = append(tests, []interface{}{0, y, "D"})
		tests = append(tests, []interface{}{len(ll) - 1, y, "U"})
	}

	fmt.Println("Part 1: ")
	fmt.Println(countFrom(ll, []interface{}{0, 0, "R"}))

	maxIlluminated := 0
	for _, test := range tests {
		count := countFrom(ll, test)
		if count > maxIlluminated {
			maxIlluminated = count
		}
	}

	fmt.Println()
	fmt.Println("Part 2: How many tiles are energized in that configuration?")
	fmt.Println(maxIlluminated)
}
