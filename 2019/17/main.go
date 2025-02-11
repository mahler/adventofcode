package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Emulator represents the IntCode computer state
type Emulator struct {
	memory       []int64
	ip           int64
	relativeBase int64
}

// NewEmulator creates a new IntCode emulator with the given program
func NewEmulator(program []int64) *Emulator {
	memory := make([]int64, len(program))
	copy(memory, program)
	return &Emulator{
		memory: memory,
	}
}

func emulate(program []int64, input <-chan int64, output chan<- int64, halt chan<- bool) {
	em := NewEmulator(program)
	defer close(halt)

	for {
		if err := em.step(input, output); err != nil {
			if err == io.EOF {
				halt <- true
				return
			}
			panic(fmt.Sprintf("emulator error: %v", err))
		}
	}
}

func (em *Emulator) getMemoryPointer(index int64) *int64 {
	// Grow memory if index is out of range
	for int64(len(em.memory)) <= index {
		em.memory = append(em.memory, 0)
	}
	return &em.memory[index]
}

func (em *Emulator) getParameter(offset, instruction int64) *int64 {
	parameter := em.memory[em.ip+offset]
	mode := instruction / pow(10, offset+1) % 10

	switch mode {
	case 0: // position mode
		return em.getMemoryPointer(parameter)
	case 1: // immediate mode
		return &parameter
	case 2: // relative mode
		return em.getMemoryPointer(em.relativeBase + parameter)
	default:
		panic(fmt.Sprintf("invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d",
			em.ip, instruction, offset, mode))
	}
}

func (em *Emulator) step(input <-chan int64, output chan<- int64) error {
	instruction := em.memory[em.ip]
	opcode := instruction % 100

	switch opcode {
	case 1: // ADD
		a, b, c := em.getParameter(1, instruction), em.getParameter(2, instruction), em.getParameter(3, instruction)
		*c = *a + *b
		em.ip += 4

	case 2: // MULTIPLY
		a, b, c := em.getParameter(1, instruction), em.getParameter(2, instruction), em.getParameter(3, instruction)
		*c = *a * *b
		em.ip += 4

	case 3: // INPUT
		a := em.getParameter(1, instruction)
		*a = <-input
		em.ip += 2

	case 4: // OUTPUT
		a := em.getParameter(1, instruction)
		output <- *a
		em.ip += 2

	case 5: // JUMP IF TRUE
		a, b := em.getParameter(1, instruction), em.getParameter(2, instruction)
		if *a != 0 {
			em.ip = *b
		} else {
			em.ip += 3
		}

	case 6: // JUMP IF FALSE
		a, b := em.getParameter(1, instruction), em.getParameter(2, instruction)
		if *a == 0 {
			em.ip = *b
		} else {
			em.ip += 3
		}

	case 7: // LESS THAN
		a, b, c := em.getParameter(1, instruction), em.getParameter(2, instruction), em.getParameter(3, instruction)
		if *a < *b {
			*c = 1
		} else {
			*c = 0
		}
		em.ip += 4

	case 8: // EQUALS
		a, b, c := em.getParameter(1, instruction), em.getParameter(2, instruction), em.getParameter(3, instruction)
		if *a == *b {
			*c = 1
		} else {
			*c = 0
		}
		em.ip += 4

	case 9: // ADJUST RELATIVE BASE
		a := em.getParameter(1, instruction)
		em.relativeBase += *a
		em.ip += 2

	case 99: // HALT
		return io.EOF

	default:
		return fmt.Errorf("invalid opcode: ip=%d instruction=%d opcode=%d", em.ip, instruction, opcode)
	}

	return nil
}

// Integer power: compute a**b using binary powering algorithm
func pow(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Direction represents a 2D vector for movement
type Direction struct {
	x, y int
}

// Common directional vectors
var (
	up    = Direction{0, -1}
	down  = Direction{0, 1}
	left  = Direction{-1, 0}
	right = Direction{1, 0}
)

// Turn maps for robot movement
var (
	turnLeft = map[Direction]Direction{
		up:    left,
		left:  down,
		down:  right,
		right: up,
	}
	turnRight = map[Direction]Direction{
		up:    right,
		right: down,
		down:  left,
		left:  up,
	}
)

// MoveList represents a sequence of movement commands
type MoveList []string

// Position represents the robot's position on the grid
type Position struct {
	pos Direction
	dir Direction
}

func main() {
	printFlag := flag.Bool("print", false, "print camera image")
	flag.Parse()

	program, err := readProgram("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read program: %v\n", err)
		os.Exit(1)
	}

	grid, err := runCameraProgram(program)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run camera program: %v\n", err)
		os.Exit(1)
	}

	if *printFlag {
		printGrid(grid)
	}

	// Part One
	fmt.Println("Part 1: What is the sum of the alignment parameters for the scaffold intersections?")
	sum := calculateAlignmentParameters(grid)
	fmt.Println(sum)

	// Part Two
	fmt.Println()
	fmt.Println("Part 2: After visiting every part of the scaffold at least once,")
	fmt.Println("how much dust does the vacuum robot report it has collected?")
	if err := solveMazePart2(program, grid); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to solve maze: %v\n", err)
		os.Exit(1)
	}
}

func readProgram(filename string) ([]int64, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var program []int64
	for _, value := range strings.Split(strings.TrimSpace(string(content)), ",") {
		num, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing int: %w", err)
		}
		program = append(program, num)
	}
	return program, nil
}

func runCameraProgram(program []int64) ([]string, error) {
	input := make(chan int64)
	output := make(chan int64)
	halt := make(chan bool)

	go emulate(program, input, output, halt)

	var builder strings.Builder
loop:
	for {
		select {
		case char := <-output:
			builder.WriteRune(rune(char))
		case <-halt:
			break loop
		}
	}

	return strings.Split(strings.TrimSpace(builder.String()), "\n"), nil
}

func calculateAlignmentParameters(grid []string) int {
	width, height := len(grid[0]), len(grid)
	sum := 0

	for y := 1; y+1 < height; y++ {
		for x := 1; x+1 < width; x++ {
			if isIntersection(grid, x, y) {
				sum += x * y
			}
		}
	}
	return sum
}

func isIntersection(grid []string, x, y int) bool {
	return grid[y][x] == '#' &&
		grid[y-1][x] == '#' &&
		grid[y+1][x] == '#' &&
		grid[y][x-1] == '#' &&
		grid[y][x+1] == '#'
}

func findRobot(grid []string) Position {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			switch grid[y][x] {
			case '^':
				return Position{Direction{x, y}, up}
			case 'v':
				return Position{Direction{x, y}, down}
			case '<':
				return Position{Direction{x, y}, left}
			case '>':
				return Position{Direction{x, y}, right}
			}
		}
	}
	return Position{}
}

func (d Direction) plus(other Direction) Direction {
	return Direction{
		x: d.x + other.x,
		y: d.y + other.y,
	}
}

func isScaffold(pos Direction, grid []string) bool {
	return pos.x >= 0 &&
		pos.y >= 0 &&
		pos.x < len(grid[0]) &&
		pos.y < len(grid) &&
		grid[pos.y][pos.x] == '#'
}

func compressPath(path MoveList, fragments []MoveList, functions []MoveList) [][4]MoveList {
	if len(functions) == 3 {
		return handleCompleteFunctions(path, fragments, functions)
	}

	if len(fragments) == 0 {
		return handleEmptyFragments(path, fragments, functions)
	}

	return handleFragmentCompression(path, fragments, functions)
}

func handleCompleteFunctions(path MoveList, fragments []MoveList, functions []MoveList) [][4]MoveList {
	if len(fragments) != 0 {
		return nil
	}

	mainFunc := createMainFunction(path, functions)
	if mainFunc == nil {
		return nil
	}

	if len(strings.Join(mainFunc, ",")) > 20 {
		return nil
	}

	return [][4]MoveList{{mainFunc, functions[0], functions[1], functions[2]}}
}

func createMainFunction(path MoveList, functions []MoveList) MoveList {
	var mainFunction MoveList
	remaining := make(MoveList, len(path))
	copy(remaining, path)

	for len(remaining) > 0 {
		found := false
		for i, function := range functions {
			if hasPrefix(remaining, function) {
				mainFunction = append(mainFunction, string('A'+i))
				remaining = remaining[len(function):]
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return mainFunction
}

func hasPrefix(list, prefix MoveList) bool {
	if len(list) < len(prefix) {
		return false
	}
	for i, move := range prefix {
		if list[i] != move {
			return false
		}
	}
	return true
}

func handleEmptyFragments(path MoveList, fragments []MoveList, functions []MoveList) [][4]MoveList {
	newFunctions := make([]MoveList, len(functions)+1)
	copy(newFunctions, functions)
	newFunctions[len(functions)] = MoveList{}
	return compressPath(path, fragments, newFunctions)
}

func handleFragmentCompression(path MoveList, fragments []MoveList, functions []MoveList) [][4]MoveList {
	var result [][4]MoveList
	fragment := fragments[0]

	for length := 1; length <= len(fragment); length++ {
		candidate := fragment[:length]
		if len(strings.Join(candidate, ",")) > 20 {
			continue
		}

		newFragments := splitFragments(fragments, candidate)
		newFunctions := append(append([]MoveList{}, functions...), candidate)

		if subresult := compressPath(path, newFragments, newFunctions); len(subresult) > 0 {
			result = append(result, subresult...)
		}
	}

	return result
}

func splitFragments(fragments []MoveList, candidate MoveList) []MoveList {
	var newFragments []MoveList
	for _, fragment := range fragments {
		current := fragment
		for {
			i := findIndex(current, candidate)
			if i == -1 {
				break
			}
			if i != 0 {
				newFragments = append(newFragments, current[:i])
			}
			current = current[i+len(candidate):]
		}
		if len(current) > 0 {
			newFragments = append(newFragments, current)
		}
	}
	return newFragments
}

func findIndex(list, sublist MoveList) int {
	for i := 0; i <= len(list)-len(sublist); i++ {
		if hasPrefix(list[i:], sublist) {
			return i
		}
	}
	return -1
}

func solveMazePart2(program []int64, grid []string) error {
	program[0] = 2 // Wake up the robot

	robotPos := findRobot(grid)
	path := generatePath(robotPos, grid)

	result := compressPath(path, []MoveList{path}, nil)
	if len(result) == 0 {
		return fmt.Errorf("no solution found")
	}

	return executeSolution(program, result[0])
}

func generatePath(pos Position, grid []string) MoveList {
	var path MoveList
	currentPos := pos

	for {
		length := 0
		for isScaffold(currentPos.pos.plus(currentPos.dir), grid) {
			currentPos.pos = currentPos.pos.plus(currentPos.dir)
			length++
		}

		if length > 0 {
			path = append(path, strconv.Itoa(length))
		}

		if newDir := turnLeft[currentPos.dir]; isScaffold(currentPos.pos.plus(newDir), grid) {
			currentPos.dir = newDir
			path = append(path, "L")
		} else if newDir := turnRight[currentPos.dir]; isScaffold(currentPos.pos.plus(newDir), grid) {
			currentPos.dir = newDir
			path = append(path, "R")
		} else {
			break
		}
	}

	return path
}

func executeSolution(program []int64, functions [4]MoveList) error {
	input := make(chan int64, 100)
	output := make(chan int64)
	halt := make(chan bool)

	go emulate(program, input, output, halt)

	commands := fmt.Sprintf("%s\n%s\n%s\n%s\nn\n",
		strings.Join(functions[0], ","),
		strings.Join(functions[1], ","),
		strings.Join(functions[2], ","),
		strings.Join(functions[3], ","))

	for _, c := range commands {
		input <- int64(c)
	}

	return processOutput(output, halt)
}

func processOutput(output <-chan int64, halt <-chan bool) error {
	for {
		select {
		case char := <-output:
			if char >= 128 {
				fmt.Println(char)
			}
		case <-halt:
			return nil
		}
	}
}

func printGrid(grid []string) {
	for _, line := range grid {
		fmt.Println(line)
	}
}
