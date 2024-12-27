package main

import (
	"fmt"
	"os"
	"strings"
)

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	R := string(data)
	G := strings.Split(R, "\n")
	H := len(G)
	W := len(G[0])

	O := make([][]int, H)
	for i := range O {
		O[i] = make([]int, W)
	}

	ax, ay := -1, -1
	for i := 0; i < H; i++ {
		if idx := strings.Index(G[i], "S"); idx != -1 {
			ax = i
			ay = idx
			break
		}
	}
	//	fmt.Println(ax, ay)

	dirs := [][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	happy := []string{"-7J", "|LJ", "-FL", "|F7"}
	Sdirs := []int{}
	for i, pos := range dirs {
		bx := ax + pos[0]
		by := ay + pos[1]
		if bx >= 0 && bx < H && by >= 0 && by < W && strings.ContainsRune(happy[i], rune(G[bx][by])) {
			Sdirs = append(Sdirs, i)
		}
	}
	//	fmt.Println(Sdirs)
	Svalid := contains(Sdirs, 3)

	transform := map[[2]interface{}]int{
		{0, "-"}: 0, {0, "7"}: 1, {0, "J"}: 3,
		{2, "-"}: 2, {2, "F"}: 1, {2, "L"}: 3,
		{1, "|"}: 1, {1, "L"}: 0, {1, "J"}: 2,
		{3, "|"}: 3, {3, "F"}: 0, {3, "7"}: 2,
	}

	curdir := Sdirs[0]
	cx := ax + dirs[curdir][0]
	cy := ay + dirs[curdir][1]
	ln := 1
	O[ax][ay] = 1
	for cx != ax || cy != ay {
		O[cx][cy] = 1
		ln++
		curdir = transform[[2]interface{}{curdir, string(G[cx][cy])}]
		cx += dirs[curdir][0]
		cy += dirs[curdir][1]
	}
	//	fmt.Println(ln)

	fmt.Println("Part 1: How many steps along the loop does it take to get from")
	fmt.Println("the starting position to the point farthest from the starting position?")
	fmt.Println(ln / 2)

	ct := 0
	for i := 0; i < H; i++ {
		inn := false
		for j := 0; j < W; j++ {
			if O[i][j] != 0 {
				if strings.ContainsRune("|JL", rune(G[i][j])) || (G[i][j] == 'S' && Svalid) {
					inn = !inn
				}
			} else {
				if inn {
					ct++
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("Part2: How many tiles are enclosed by the loop?")
	fmt.Println(ct)
}
