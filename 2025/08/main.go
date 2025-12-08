package main

import (
        "bufio"
        "fmt"
        "math"
        "os"
        "sort"
        "strconv"
        "strings"
)

const filename = "input.txt"

type Point []int

type Distance struct {
        First  Point
        Second Point
        Dist   float64
}

func parseInput(filename string) ([]Point, error) {
        file, err := os.Open(filename)
        if err != nil {
                return nil, err
        }
        defer file.Close()

        var coordinates []Point
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := strings.TrimSpace(scanner.Text())
                if line == "" {
                        continue
                }
                parts := strings.Split(line, ",")
                point := make(Point, len(parts))
                for i, numStr := range parts {
                        num, err := strconv.Atoi(strings.TrimSpace(numStr))
                        if err != nil {
                                return nil, err
                        }
                        point[i] = num
                }
                coordinates = append(coordinates, point)
        }
        return coordinates, scanner.Err()
}

func distance(a, b Point) float64 {
        sum := 0.0
        for i := range a {
                diff := float64(a[i] - b[i])
                sum += diff * diff
        }
        return math.Sqrt(sum)
}

func sortDistances(coordinates []Point) []Distance {
        var distances []Distance
        for i, a := range coordinates {
                for _, b := range coordinates[i+1:] {
                        distances = append(distances, Distance{
                                First:  a,
                                Second: b,
                                Dist:   distance(a, b),
                        })
                }
        }
        sort.Slice(distances, func(i, j int) bool {
                return distances[i].Dist < distances[j].Dist
        })
        return distances
}

func pointKey(p Point) string {
        parts := make([]string, len(p))
        for i, v := range p {
                parts[i] = strconv.Itoa(v)
        }
        return strings.Join(parts, ",")
}

func partOneScore(circuits map[string]int) int {
        counter := make(map[int]int)
        for _, v := range circuits {
                counter[v]++
        }

        var counts []int
        for _, count := range counter {
                counts = append(counts, count)
        }
        sort.Slice(counts, func(i, j int) bool {
                return counts[i] > counts[j]
        })

        product := 1
        for i := 0; i < 3 && i < len(counts); i++ {
                product *= counts[i]
        }
        return product
}

func solve(distances []Distance, boxCount, connectionLimit int) (int, int) {
        var p1, p2 int
        circuits := make(map[string]int)
        counter := 0

        for i, d := range distances {
                if i == connectionLimit {
                        p1 = partOneScore(circuits)
                }

                firstKey := pointKey(d.First)
                secondKey := pointKey(d.Second)

                g1, hasG1 := circuits[firstKey]
                g2, hasG2 := circuits[secondKey]

                if hasG1 && hasG2 {
                        if g1 == g2 {
                                continue
                        }
                        // Merge g2 into g1
                        for k, v := range circuits {
                                if v == g2 {
                                        circuits[k] = g1
                                }
                        }
                } else if !hasG1 && !hasG2 {
                        circuits[firstKey] = counter
                        circuits[secondKey] = counter
                        counter++
                } else if !hasG1 {
                        circuits[firstKey] = g2
                } else {
                        circuits[secondKey] = g1
                }

                // Check if all boxes are in one circuit
                if len(circuits) == boxCount {
                        uniqueCircuits := make(map[int]bool)
                        for _, v := range circuits {
                                uniqueCircuits[v] = true
                        }
                        if len(uniqueCircuits) == 1 {
                                p2 = d.First[0] * d.Second[0]
                                break
                        }
                }
        }

        return p1, p2
}

func main() {
        coordinates, err := parseInput("input.txt")
        if err != nil {
                fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
                os.Exit(1)
        }

        boxCount := len(coordinates)
        distances := sortDistances(coordinates)
        partOne, partTwo := solve(distances, boxCount, 1000)

        fmt.Println(partOne)
        fmt.Println(partTwo)
}
