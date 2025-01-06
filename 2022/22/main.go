package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	D      = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	REGION = [][2]int{{0, 1}, {0, 2}, {1, 1}, {2, 1}, {2, 0}, {3, 0}}
)

func regionToGlobal(r, c, region, cube int) (int, int) {
	rr, cc := REGION[region-1][0], REGION[region-1][1]
	return rr*cube + r, cc*cube + c
}

func getRegion(r, c, cube int) (int, int, int) {
	for i, rc := range REGION {
		rr, cc := rc[0], rc[1]
		if rr*cube <= r && r < (rr+1)*cube && cc*cube <= c && c < (cc+1)*cube {
			return i + 1, r - rr*cube, c - cc*cube
		}
	}
	panic(fmt.Sprintf("Invalid region: r=%d, c=%d", r, c))
}

func newCoords(r, c, d, nd, cube int) (int, int) {
	var x int
	switch d {
	case 0:
		x = c
	case 1:
		x = r
	case 2:
		x = cube - 1 - c
	case 3:
		x = cube - 1 - r
	}

	switch nd {
	case 0:
		return cube - 1, x
	case 1:
		return x, 0
	case 2:
		return 0, cube - 1 - x
	case 3:
		return cube - 1 - x, cube - 1
	}
	panic("Invalid direction")
}

func getDest(G []string, r, c, d, part, R, C, cube int) (int, int, int) {
	if part == 1 {
		rr := (r + D[d][0] + R) % R
		cc := (c + D[d][1] + C) % C
		for G[rr][cc] == ' ' {
			rr = (rr + D[d][0] + R) % R
			cc = (cc + D[d][1] + C) % C
		}
		return rr, cc, d
	}

	region, rr, rc := getRegion(r, c, cube)
	transitions := map[[2]int][2]int{
		{4, 0}: {3, 0}, {4, 1}: {2, 3}, {4, 2}: {6, 3}, {4, 3}: {5, 3},
		{1, 0}: {6, 1}, {1, 1}: {2, 1}, {1, 2}: {3, 2}, {1, 3}: {5, 1},
		{3, 0}: {1, 0}, {3, 1}: {2, 0}, {3, 2}: {4, 2}, {3, 3}: {5, 2},
		{6, 0}: {5, 0}, {6, 1}: {4, 0}, {6, 2}: {2, 2}, {6, 3}: {1, 2},
		{2, 0}: {6, 0}, {2, 1}: {4, 3}, {2, 2}: {3, 3}, {2, 3}: {1, 3},
		{5, 0}: {3, 1}, {5, 1}: {4, 1}, {5, 2}: {6, 2}, {5, 3}: {1, 1},
	}
	ndRegion := transitions[[2]int{region, d}]
	newRegion, nd := ndRegion[0], ndRegion[1]
	nr, nc := newCoords(rr, rc, d, nd, cube)
	nr, nc = regionToGlobal(nr, nc, newRegion, cube)
	if G[nr][nc] != '.' && G[nr][nc] != '#' {
		panic(fmt.Sprintf("Invalid destination: %c", G[nr][nc]))
	}
	return nr, nc, nd
}

func solve(G []string, instr string, R, C, cube, part int) int {
	r, c, d := 0, 0, 1
	for G[r][c] != '.' {
		c++
	}

	i := 0
	for i < len(instr) {
		n := 0
		for i < len(instr) && instr[i] >= '0' && instr[i] <= '9' {
			n = n*10 + int(instr[i]-'0')
			i++
		}
		for step := 0; step < n; step++ {
			rr, cc := (r+D[d][0]+R)%R, (c+D[d][1]+C)%C
			if G[rr][cc] == ' ' {
				nr, nc, nd := getDest(G, r, c, d, part, R, C, cube)
				if G[nr][nc] == '#' {
					break
				}
				r, c, d = nr, nc, nd
			} else if G[rr][cc] == '#' {
				break
			} else {
				r, c = rr, cc
			}
		}
		if i == len(instr) {
			break
		}
		turn := instr[i]
		if turn == 'L' {
			d = (d + 3) % 4
		} else if turn == 'R' {
			d = (d + 1) % 4
		} else {
			panic("Invalid instruction")
		}
		i++
	}

	DV := map[int]int{0: 3, 1: 0, 2: 1, 3: 2}
	return (r+1)*1000 + (c+1)*4 + DV[d]
}

func main() {
	// Read input file
	infile := "input.txt"
	data, err := os.ReadFile(infile)
	if err != nil {
		panic(err)
	}

	// Parse input
	content := string(data)
	parts := strings.Split(content, "\n\n")
	G := strings.Split(parts[0], "\n")
	instr := strings.TrimSpace(parts[1])

	R := len(G)
	C := len(G[0])
	for i := range G {
		for len(G[i]) < C {
			G[i] += " "
		}
	}

	CUBE := C / 3
	if CUBE != R/4 {
		panic("Invalid cube dimensions")
	}

	fmt.Println("Part 1: Follow the path given in the monkeys' notes. What is the final password?")
	fmt.Println(solve(G, instr, R, C, CUBE, 1))

	fmt.Println()
	fmt.Println("Part 2: Fold the map into a cube, then follow the path given in the monkeys' notes.")
	fmt.Println("What is the final password?")
	fmt.Println(solve(G, instr, R, C, CUBE, 2))
}
