package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game represents the Intcode computer state
type Game struct {
	prog    map[int]int
	ip      int
	input   []int
	output  []int
	relBase int
	halt    bool
}

// NewGame creates a new Game instance
func NewGame(prog map[int]int) *Game {
	return &Game{
		prog:    prog,
		ip:      0,
		input:   make([]int, 0),
		output:  make([]int, 0),
		relBase: 0,
		halt:    false,
	}
}

// splitInstruction splits an instruction into opcode and modes
func splitInstruction(instruction int) (string, string) {
	instrStr := fmt.Sprintf("%05d", instruction)
	return instrStr[3:], instrStr[:3]
}

// getValues gets the parameter values based on modes
func getValues(input map[int]int, pos int, op string, modes string, game *Game) ([]int, int) {
	modeC := string(modes[2])
	modeB := string(modes[1])
	modeA := string(modes[0])
	values := make([]int, 0)
	offset := 0

	if strings.Contains("01,02,04,05,06,07,08,09", op) {
		switch modeC {
		case "0":
			values = append(values, input[input[pos+1]])
		case "1":
			values = append(values, input[pos+1])
		case "2":
			values = append(values, input[input[pos+1]+game.relBase])
		}

		if strings.Contains("01,02,05,06,07,08", op) {
			switch modeB {
			case "0":
				values = append(values, input[input[pos+2]])
			case "1":
				values = append(values, input[pos+2])
			case "2":
				values = append(values, input[input[pos+2]+game.relBase])
			}
		}
	}

	if strings.Contains("01,02,07,08", op) && modeA == "2" {
		offset = game.relBase
	}

	if op == "03" && modeC == "2" {
		offset = game.relBase
	}

	return values, offset
}

// readOutput reads and removes the first output value
func readOutput(game *Game) (int, bool) {
	if len(game.output) > 0 {
		output := game.output[0]
		game.output = game.output[1:]
		return output, true
	}
	return 0, false
}

// runGame runs the Intcode program
func runGame(game *Game, input []int) {
	if len(input) > 0 {
		game.input = append(game.input, input...)
	}

	for game.prog[game.ip] != 99 {
		op, modes := splitInstruction(game.prog[game.ip])
		values, offset := getValues(game.prog, game.ip, op, modes, game)

		switch op {
		case "01": // Addition
			game.prog[game.prog[game.ip+3]+offset] = values[0] + values[1]
			game.ip += 4

		case "02": // Multiplication
			game.prog[game.prog[game.ip+3]+offset] = values[0] * values[1]
			game.ip += 4

		case "03": // Read and Store input
			if len(game.input) == 0 {
				return
			}
			game.prog[game.prog[game.ip+1]+offset] = game.input[0]
			game.input = game.input[1:]
			game.ip += 2

		case "04": // Print Output
			game.output = append(game.output, values[0])
			game.ip += 2
			return

		case "05": // Jump-if-True
			if values[0] != 0 {
				game.ip = values[1]
			} else {
				game.ip += 3
			}

		case "06": // Jump-if-False
			if values[0] == 0 {
				game.ip = values[1]
			} else {
				game.ip += 3
			}

		case "07": // Less than
			if values[0] < values[1] {
				game.prog[game.prog[game.ip+3]+offset] = 1
			} else {
				game.prog[game.prog[game.ip+3]+offset] = 0
			}
			game.ip += 4

		case "08": // Equals
			if values[0] == values[1] {
				game.prog[game.prog[game.ip+3]+offset] = 1
			} else {
				game.prog[game.prog[game.ip+3]+offset] = 0
			}
			game.ip += 4

		case "09": // Adjust Relative Base
			game.relBase += values[0]
			game.ip += 2
		}
	}

	game.halt = true
}

// createProgram creates the initial program state
func createProgram(input []string) map[int]int {
	prog := make(map[int]int)
	for i, val := range input {
		num, _ := strconv.Atoi(val)
		prog[i] = num
	}
	return prog
}

// readFile reads the input file
func readFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines
}

func runWalkInstruction(game *Game) {
	output := 0
	var ok bool

	for output != 10 {
		runGame(game, nil)
		if output, ok = readOutput(game); ok {
			//fmt.Print(string(output))
		}
	}

	instruction := "NOT C J\n" +
		"NOT A T\n" +
		"OR T J\n" +
		"AND D J\n" +
		"WALK\n"

	//fmt.Print("\nExecuting WALK instruction:\n")
	//fmt.Print(instruction)

	inputChars := make([]int, 0)
	for _, char := range instruction {
		inputChars = append(inputChars, int(char))
	}

	runGame(game, inputChars)
	fmt.Println("Part 1: What amount of hull damage does it report?")
	processOutput(game)
}

func runRunInstruction(game *Game) {
	output := 0
	var ok bool

	for output != 10 {
		runGame(game, nil)
		if output, ok = readOutput(game); ok {
			//fmt.Print(string(output))
		}
	}

	instruction := "NOT C J\n" +
		"NOT B T\n" +
		"OR T J\n" +
		"NOT A T\n" +
		"OR T J\n" +
		"AND D J\n" +
		"NOT E T\n" +
		"NOT T T\n" +
		"OR H T\n" +
		"AND T J\n" +
		"RUN\n"

	//fmt.Print("\nExecuting RUN instruction:\n")
	//fmt.Print(instruction)

	inputChars := make([]int, 0)
	for _, char := range instruction {
		inputChars = append(inputChars, int(char))
	}

	runGame(game, inputChars)
	fmt.Println()
	fmt.Println("Part 2: What amount of hull damage does the springdroid now report?")
	processOutput(game)
}

func processOutput(game *Game) {
	var output int
	var ok bool

	for !game.halt {
		if output, ok = readOutput(game); ok {
			if output > 255 {
				fmt.Println(output)
				return
			}
			//fmt.Print(string(output))
		}
		runGame(game, nil)
	}
}

func main() {
	input := strings.Split(readFile()[0], ",")

	// Run with WALK instruction
	progWalk := createProgram(input)
	gameWalk := NewGame(progWalk)
	runWalkInstruction(gameWalk)

	// Run with RUN instruction
	progRun := createProgram(input)
	gameRun := NewGame(progRun)
	runRunInstruction(gameRun)
}
