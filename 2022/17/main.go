package main

import (
	"fmt"
	"os"
	"strings"
)

type Complex struct {
	real, imag float64
}

func (c Complex) add(other Complex) Complex {
	return Complex{c.real + other.real, c.imag + other.imag}
}

type Rock []Complex
type Tower map[Complex]bool

func main() {
	rocks := []Rock{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
	}

	data, _ := os.ReadFile("input.txt")
	jets := make([]int, 0)
	for _, c := range strings.TrimSpace(string(data)) {
		jets = append(jets, int(c)-61)
	}

	tower := make(Tower)
	cache := make(map[string][]int)
	top := 0
	i, j := 0, 0

	empty := func(pos Complex) bool {
		return pos.real >= 0 && pos.real < 7 && pos.imag > 0 && !tower[pos]
	}

	check := func(pos Complex, dir Complex, rock Rock) bool {
		for _, r := range rock {
			newPos := pos.add(dir).add(r)
			if !empty(newPos) {
				return false
			}
		}
		return true
	}

	target := int64(1e12)
	for step := int64(0); step < target; step++ {
		pos := Complex{2, float64(top + 4)}

		if step == 2022 {
			// We're there now...
			fmt.Println("Part 1: How many units tall will the tower of rocks be after 2022 rocks have stopped falling?")
			fmt.Println(top)
		}

		key := fmt.Sprintf("%d,%d", i, j)
		if val, exists := cache[key]; exists {
			s, t := val[0], val[1]
			d, m := (target-step)/int64(step-int64(s)), (target-step)%int64(step-int64(s))
			if m == 0 {
				fmt.Println()
				fmt.Println("Part 2: How tall will the tower be after 1000000000000 rocks have stopped?")
				fmt.Printf("%.0f\n", float64(top)+float64(top-t)*float64(d))
				break
			}
		} else {
			cache[key] = []int{int(step), top}
		}

		rock := rocks[i]
		i = (i + 1) % len(rocks)

		for {
			jet := Complex{float64(jets[j]), 0}
			j = (j + 1) % len(jets)

			if check(pos, jet, rock) {
				pos = pos.add(jet)
			}
			if check(pos, Complex{0, -1}, rock) {
				pos = pos.add(Complex{0, -1})
			} else {
				break
			}
		}

		heights := []int{1, 0, 2, 2, 3}
		for _, r := range rock {
			tower[pos.add(r)] = true
			if int(pos.imag)+heights[i] > top {
				top = int(pos.imag) + heights[i]
			}
		}
	}
}
