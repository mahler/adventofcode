package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cart struct {
	x, y         int
	dx, dy       int
	intersection int
	crashed      bool
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the track layout
	var track [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		track = append(track, []byte(scanner.Text()))
	}

	// Find carts and initialize them
	var carts []Cart
	for y := 0; y < len(track); y++ {
		for x := 0; x < len(track[y]); x++ {
			var cart Cart
			switch track[y][x] {
			case '<':
				cart = Cart{x: x, y: y, dx: -1, dy: 0}
				track[y][x] = '-'
			case '>':
				cart = Cart{x: x, y: y, dx: 1, dy: 0}
				track[y][x] = '-'
			case '^':
				cart = Cart{x: x, y: y, dx: 0, dy: -1}
				track[y][x] = '|'
			case 'v':
				cart = Cart{x: x, y: y, dx: 0, dy: 1}
				track[y][x] = '|'
			default:
				continue
			}
			carts = append(carts, cart)
		}
	}

	firstCrash := true
	for {
		// Check if only one cart remains
		activeCarts := 0
		for _, cart := range carts {
			if !cart.crashed {
				activeCarts++
			}
		}
		if activeCarts == 1 {
			break
		}

		// Sort carts by position (top to bottom, left to right)
		order := make([]int, len(carts))
		idx := 0
		for y := 0; y < len(track); y++ {
			for x := 0; x < len(track[0]); x++ {
				for i := range carts {
					if !carts[i].crashed && carts[i].x == x && carts[i].y == y {
						order[idx] = i
						idx++
					}
				}
			}
		}
		for i := range carts {
			if carts[i].crashed {
				order[idx] = i
				idx++
			}
		}

		// Move carts
		for _, i := range order {
			if carts[i].crashed {
				continue
			}

			// Move cart
			carts[i].x += carts[i].dx
			carts[i].y += carts[i].dy

			// Handle track types
			switch track[carts[i].y][carts[i].x] {
			case '/':
				carts[i].dx, carts[i].dy = -carts[i].dy, -carts[i].dx
			case '\\':
				carts[i].dx, carts[i].dy = carts[i].dy, carts[i].dx
			case '+':
				switch carts[i].intersection {
				case 0: // Turn left
					carts[i].dx, carts[i].dy = carts[i].dy, -carts[i].dx
					carts[i].intersection = 1
				case 1: // Go straight
					carts[i].intersection = 2
				case 2: // Turn right
					carts[i].dx, carts[i].dy = -carts[i].dy, carts[i].dx
					carts[i].intersection = 0
				}
			}

			// Check for collisions
			for j := range carts {
				if j != i && !carts[j].crashed && carts[j].x == carts[i].x && carts[j].y == carts[i].y {
					if firstCrash {
						fmt.Println("Part 1: To help prevent crashes, you'd like to know the location of the first crash...")
						fmt.Printf("%d,%d\n", carts[i].x, carts[i].y)
						firstCrash = false
					}
					carts[j].crashed = true
					carts[i].crashed = true
					break
				}
			}
		}
	}

	// Find the last remaining cart
	for _, cart := range carts {
		if !cart.crashed {
			fmt.Println()
			fmt.Println("Part 2: What is the location of the last cart at the end of the first tick where it is the only cart left?")
			fmt.Printf("%d,%d\n", cart.x, cart.y)
			break
		}
	}
}
