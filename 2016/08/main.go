package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 8: Two-Factor Authentication")
	fileContent, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileRows := strings.Split(string(fileContent), "\n")

	rectangleInstruction := regexp.MustCompile(`rect (\d+)x(\d+)`)
	rotateRowInstruction := regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
	rotateColInstruction := regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)

	width := 50
	height := 6

	data := make([]bool, width*height)

	for _, instruction := range fileRows {
		if result := rectangleInstruction.FindStringSubmatch(instruction); result != nil {
			w, _ := strconv.Atoi(result[1])
			h, _ := strconv.Atoi(result[2])
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					data[y*width+x] = true
				}
			}
		} else if result := rotateRowInstruction.FindStringSubmatch(instruction); result != nil {
			y, _ := strconv.Atoi(result[1])
			a, _ := strconv.Atoi(result[2])
			for i := 0; i < a; i++ {
				tmp := data[y*width+width-1]
				for x := width - 1; x > 0; x-- {
					data[y*width+x] = data[y*width+x-1]
				}
				data[y*width+0] = tmp
			}
		} else if result := rotateColInstruction.FindStringSubmatch(instruction); result != nil {
			x, _ := strconv.Atoi(result[1])
			a, _ := strconv.Atoi(result[2])
			for i := 0; i < a; i++ {
				tmp := data[(height-1)*width+x]
				for y := height - 1; y > 0; y-- {
					data[y*width+x] = data[(y-1)*width+x]
				}
				data[0*width+x] = tmp
			}
		} else {
			panic(instruction)
		}
	}

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y*width+x] {
				count++
			}
		}
	}
	fmt.Println("how many pixels should be lit?")
	fmt.Println(count)

	fmt.Println()
	fmt.Println("Part 2/")
	fmt.Println("what code is the screen trying to display?")
	fmt.Println()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y*width+x] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
