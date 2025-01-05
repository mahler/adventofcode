package main

import (
	"bytes"
	"fmt"
	"os"
)

type Cube struct {
	x int
	y int
	z int
}

func markExterior(cube Cube, cubeSpace map[Cube]bool, exterior map[Cube]bool, cubeSize Cube) int {
	if exterior[cube] {
		return 0
	}
	if cube.x < -1 || cube.x > cubeSize.x+1 || cube.y < -1 || cube.y > cubeSize.y+1 || cube.z < -1 || cube.z > cubeSize.z+1 {
		return 0
	}
	if cubeSpace[cube] {
		return 1
	}
	exterior[cube] = true
	var count int
	for _, neighbor := range neighbors(cube) {
		count += markExterior(neighbor, cubeSpace, exterior, cubeSize)
	}
	return count
}

func neighbors(cube Cube) []Cube {
	return []Cube{
		{x: cube.x - 1, y: cube.y, z: cube.z},
		{x: cube.x + 1, y: cube.y, z: cube.z},
		{x: cube.x, y: cube.y - 1, z: cube.z},
		{x: cube.x, y: cube.y + 1, z: cube.z},
		{x: cube.x, y: cube.y, z: cube.z - 1},
		{x: cube.x, y: cube.y, z: cube.z + 1},
	}
}

func processCube(input []byte) []Cube {
	var cubes []Cube
	for _, line := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		coordinates := bytes.Split(bytes.TrimSpace(line), []byte{','})
		if len(coordinates) != 3 {
			panic("invalid input")
		}
		cubes = append(cubes, Cube{
			x: atoi(coordinates[0]),
			y: atoi(coordinates[1]),
			z: atoi(coordinates[2]),
		})
	}
	return cubes
}

func atoi(b []byte) int {
	var n int
	for _, c := range b {
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// part1
	cubes := processCube(input)
	cubeSpace := make(map[Cube]bool)
	for _, cube := range cubes {
		cubeSpace[cube] = true
	}
	var totalSurfaceArea int
	for _, cube := range cubes {
		var surfaceArea int
		for _, neighbor := range neighbors(cube) {
			if !cubeSpace[neighbor] {
				surfaceArea++
			}
		}
		totalSurfaceArea += surfaceArea
	}
	fmt.Println("Part 1: What is the surface area of your scanned lava droplet?")
	fmt.Println(totalSurfaceArea)

	//part2
	cubeSize := Cube{}
	for _, cube := range cubes {
		if cube.x > cubeSize.x {
			cubeSize.x = cube.x
		}
		if cube.y > cubeSize.y {
			cubeSize.y = cube.y
		}
		if cube.z > cubeSize.z {
			cubeSize.z = cube.z
		}
	}

	for _, cube := range cubes {
		cubeSpace[cube] = true
	}
	exterior := make(map[Cube]bool)
	count := markExterior(Cube{x: 0, y: 0, z: 0}, cubeSpace, exterior, cubeSize)

	fmt.Println()
	fmt.Println("Part 2: What is the exterior surface area of your scanned lava droplet?")
	fmt.Println(count)
}
