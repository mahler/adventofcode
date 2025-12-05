package main

import (
        "bufio"
        "fmt"
        "os"
)

type Point struct {
        x, y int
}

func main() {
        file, err := os.Open("input.txt")
        if err != nil {
                fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
                os.Exit(1)
        }
        defer file.Close()

        // Read warehouse and find rolls
        rolls := make(map[Point]bool)
        scanner := bufio.NewScanner(file)
        y := 0
        for scanner.Scan() {
                line := scanner.Text()
                for x, bin := range line {
                        if bin == '@' {
                                rolls[Point{x, y}] = true
                        }
                }
                y++
        }

        if err := scanner.Err(); err != nil {
                fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
                os.Exit(1)
        }

        adjacents := []Point{
                {-1, -1}, {0, -1}, {1, -1},
                {-1, 0}, {1, 0},
                {-1, 1}, {0, 1}, {1, 1},
        }

        a1, a2, round := 0, 0, 0
        elvesWorking := true
        removed := make(map[Point]bool)

        for elvesWorking {
                round++
                for roll := range rolls {
                        adjacentRolls := 0
                        for _, adj := range adjacents {
                                neighbor := Point{roll.x + adj.x, roll.y + adj.y}
                                if rolls[neighbor] {
                                        adjacentRolls++
                                }
                                if adjacentRolls > 3 {
                                        break
                                }
                        }
                        if adjacentRolls < 4 {
                                removed[roll] = true
                        }
                }

                if len(removed) > 0 {
                        // Remove from rolls
                        for r := range removed {
                                delete(rolls, r)
                        }

                        if round == 1 {
                                a1 += len(removed)
                        }
                        a2 += len(removed)

                        // Clear removed set
                        removed = make(map[Point]bool)
                } else {
                        elvesWorking = false
                }
        }

        fmt.Printf("Part one: %d rolls removed\n", a1)
        fmt.Printf("Part two: %d rolls removed in %d rounds of work\n", a2, round)
}
