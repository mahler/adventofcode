package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
        "strings"
)

func part1() {
        file, err := os.Open("input.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        currentPosition := 50
        numberOfZeros := 0

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                cleanValue := strings.TrimSpace(line)

                if len(cleanValue) == 0 {
                        continue
                }

                direction := cleanValue[0]
                distance, err := strconv.Atoi(cleanValue[1:])
                if err != nil {
                        log.Fatalf("Invalid distance value: %v", err)
                }

                switch direction {
                case 'L':
                        currentPosition -= distance
                case 'R':
                        currentPosition += distance
                default:
                        log.Fatal("Invalid direction character. Must be 'L' or 'R'.")
                }

                currentPosition = ((currentPosition % 100) + 100) % 100
                if currentPosition == 0 {
                        numberOfZeros++
                }

                // fmt.Printf("Direction: %c, Distance: %d\n", direction, distance)
                // fmt.Printf("Current Position: %d\n", currentPosition)
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        fmt.Printf("Number of times position was zero: %d\n", numberOfZeros)
}

func part2() {
        file, err := os.Open("input.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        currentPosition := 50
        numberOfZeros := 0

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                cleanValue := strings.TrimSpace(line)

                if len(cleanValue) == 0 {
                        continue
                }

                direction := cleanValue[0]
                distance, err := strconv.Atoi(cleanValue[1:])
                if err != nil {
                        log.Fatalf("Invalid distance value: %v", err)
                }

                var newPosition int
                switch direction {
                case 'L':
                        newPosition = ((currentPosition - distance) % 100 + 100) % 100
                        if newPosition > currentPosition && currentPosition != 0 {
                                numberOfZeros++
                        }
                case 'R':
                        newPosition = (currentPosition + distance) % 100
                        if newPosition < currentPosition && newPosition != 0 {
                                numberOfZeros++
                        }
                default:
                        log.Fatal("Invalid direction character. Must be 'L' or 'R'.")
                }

                currentPosition = newPosition

                // land directly on zero
                if currentPosition == 0 {
                        numberOfZeros++
                }

                // cross over a zero
                numberOfRotations := distance / 100
                numberOfZeros += numberOfRotations

                // fmt.Printf("Direction: %c, Distance: %d\n", direction, distance)
                // fmt.Printf("Current Position: %d\n", currentPosition)
                // fmt.Printf("Number of times position was zero: %d\n", numberOfZeros)
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        fmt.Printf("Number of times position was zero: %d\n", numberOfZeros)
}

func main() {
        part1()
        part2()
}
