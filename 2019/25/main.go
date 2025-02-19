package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParameterMode int

const (
	Position  ParameterMode = 0
	Immediate ParameterMode = 1
	Relative  ParameterMode = 2
)

type Intcode struct {
	verbose       bool
	program       []int
	inputStream   []int
	ip            int
	relBase       int
	eop           bool
	needsInput    bool
	readOutput    bool
	stallOnOutput bool
	retVal        int
}

type opcodeFunc func([]int)

func NewIntcode(verbose bool) *Intcode {
	return &Intcode{
		verbose: verbose,
	}
}

func (ic *Intcode) printOperandValues(numOperands int) string {
	if numOperands > 0 {
		operands := make([]string, numOperands)
		for i := 0; i < numOperands; i++ {
			operands[i] = strconv.Itoa(ic.program[ic.ip+1+i])
		}
		return strings.Join(operands, ",")
	}
	return ""
}

func (ic *Intcode) initProgram(program []int, inputStream []int, stallOnOutput bool) {
	ic.program = make([]int, len(program)+5000)
	copy(ic.program, program)
	ic.ip = 0
	ic.relBase = 0
	ic.eop = false
	ic.needsInput = false
	ic.readOutput = false
	ic.stallOnOutput = stallOnOutput
	ic.retVal = 0
	ic.inputStream = inputStream
}

func (ic *Intcode) runProgram(inputStream []int) int {
	ic.readOutput = false
	if inputStream != nil {
		ic.inputStream = append(ic.inputStream, inputStream...)
		ic.needsInput = false
	}

	for !ic.eop && !ic.needsInput && (!ic.stallOnOutput || !ic.readOutput) {
		instruction := fmt.Sprintf("%014d%d", 0, ic.program[ic.ip])
		opcode := instruction[len(instruction)-2:]
		opcodeInt, _ := strconv.Atoi(opcode)

		var op opcodeFunc
		var params int
		var name string

		switch opcodeInt {
		case 1:
			op, params, name = ic.opcode1, 3, "add"
		case 2:
			op, params, name = ic.opcode2, 3, "mul"
		case 3:
			op, params, name = ic.opcode3, 1, "inp"
		case 4:
			op, params, name = ic.opcode4, 1, "out"
		case 5:
			op, params, name = ic.opcode5, 2, "jmp1"
		case 6:
			op, params, name = ic.opcode6, 2, "jmp0"
		case 7:
			op, params, name = ic.opcode7, 3, "lt"
		case 8:
			op, params, name = ic.opcode8, 3, "eq"
		case 9:
			op, params, name = ic.opcode9, 1, "rel"
		case 99:
			op, params, name = ic.opcode99, 0, "eop"
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", opcodeInt))
		}

		if ic.verbose {
			fmt.Printf("IP: %d - %s(%s) - Operands(%s) - \n", ic.ip, instruction, name, ic.printOperandValues(params))
		}

		modes := make([]int, len(instruction)-2)
		for i := range modes {
			modes[i], _ = strconv.Atoi(string(instruction[len(instruction)-3-i]))
		}

		op(modes)
	}

	return ic.retVal
}

func (ic *Intcode) getOperand(operand int, modes []int, isDest bool) int {
	value := ic.program[ic.ip+operand]
	mode := ParameterMode(modes[operand-1])

	if isDest {
		if mode == Position {
			if ic.verbose {
				fmt.Printf("    (Operand %d, Mode %d) = DST(%d)\n", operand, mode, value)
			}
			return value
		} else if mode == Relative {
			if ic.verbose {
				fmt.Printf("    (Operand %d, Mode %d, RelBase %d, Value %d) = DST(%d)\n",
					operand, mode, ic.relBase, value, ic.relBase+value)
			}
			return ic.relBase + value
		}
	} else {
		if mode == Position {
			if ic.verbose {
				fmt.Printf("    (Operand %d, Mode %d) = %d\n", operand, mode, ic.program[value])
			}
			return ic.program[value]
		} else if mode == Immediate {
			if ic.verbose {
				fmt.Printf("    (Operand %d, Mode %d) = %d\n", operand, mode, value)
			}
			return value
		} else if mode == Relative {
			if ic.verbose {
				fmt.Printf("    (Operand %d, Mode %d, RelBase %d, Value %d) = %d\n",
					operand, mode, ic.relBase, value, ic.program[ic.relBase+value])
			}
			return ic.program[ic.relBase+value]
		}
	}
	return 0
}

func (ic *Intcode) opcode1(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	src1 := ic.getOperand(2, modes, false)
	dest := ic.getOperand(3, modes, true)
	ic.program[dest] = src0 + src1
	ic.ip += 4
}

func (ic *Intcode) opcode2(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	src1 := ic.getOperand(2, modes, false)
	dest := ic.getOperand(3, modes, true)
	ic.program[dest] = src0 * src1
	ic.ip += 4
}

func (ic *Intcode) opcode3(modes []int) {
	if len(ic.inputStream) == 0 {
		ic.needsInput = true
	} else {
		dest := ic.getOperand(1, modes, true)
		value := ic.inputStream[0]
		ic.inputStream = ic.inputStream[1:]
		ic.program[dest] = value
		ic.ip += 2
	}
}

func (ic *Intcode) opcode4(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	if ic.verbose {
		fmt.Printf("Output: %d\n", src0)
	}
	ic.retVal = src0
	ic.readOutput = true
	ic.ip += 2
}

func (ic *Intcode) opcode5(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	jumpLoc := ic.getOperand(2, modes, false)
	if src0 != 0 {
		ic.ip = jumpLoc
	} else {
		ic.ip += 3
	}
}

func (ic *Intcode) opcode6(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	jumpLoc := ic.getOperand(2, modes, false)
	if src0 == 0 {
		ic.ip = jumpLoc
	} else {
		ic.ip += 3
	}
}

func (ic *Intcode) opcode7(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	src1 := ic.getOperand(2, modes, false)
	dest := ic.getOperand(3, modes, true)
	if src0 < src1 {
		ic.program[dest] = 1
	} else {
		ic.program[dest] = 0
	}
	ic.ip += 4
}

func (ic *Intcode) opcode8(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	src1 := ic.getOperand(2, modes, false)
	dest := ic.getOperand(3, modes, true)
	if src0 == src1 {
		ic.program[dest] = 1
	} else {
		ic.program[dest] = 0
	}
	ic.ip += 4
}

func (ic *Intcode) opcode9(modes []int) {
	src0 := ic.getOperand(1, modes, false)
	ic.relBase += src0
	ic.ip += 2
}

func (ic *Intcode) opcode99(modes []int) {
	ic.ip += 1
	ic.eop = true
	if ic.retVal == 0 {
		ic.retVal = ic.program[0]
	}
}

// ---- Actual program

func main() {
	// Commands acts like a save game and replay the commands in sequence when the program launch
	// The "ready to test items" is a brute force to find the four times needed to solve task.
	commands := []string{
		"north",
		"west",
		"take planetoid",
		"west",
		"take spool of cat6",
		"east",
		"east",
		"south",
		"west",
		"north",
		"take dark matter",
		"south",
		"east",
		"east",
		"north",
		"take sand",
		"west",
		"take coin",
		"west",
		"south",
		"west",
		"take fuel cell",
		"east",
		"take wreath",
		"north",
		"east",
		"north",
		"take jam",
		"south",
		"west",
		"north",
		"west",
		"drop jam",
		"drop fuel cell",
		"drop planetoid",
		"drop sand",
		"drop spool of cat6",
		"drop coin",
		"drop dark matter",
		"drop wreath",
		// ready to test items
		"take jam", "take fuel cell", "take planetoid", "take sand",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take planetoid", "take spool of cat6",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take planetoid", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take planetoid", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take planetoid", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take sand", "take spool of cat6",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take sand", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take sand", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take sand", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take spool of cat6", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take spool of cat6", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take spool of cat6", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take fuel cell", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take sand", "take spool of cat6",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take sand", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take sand", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take sand", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take spool of cat6", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take spool of cat6", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take spool of cat6", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take planetoid", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take sand", "take spool of cat6", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take sand", "take spool of cat6", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take sand", "take spool of cat6", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take sand", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take sand", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take sand", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take spool of cat6", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take spool of cat6", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take spool of cat6", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take jam", "take coin", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take sand", "take spool of cat6",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take sand", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take sand", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take sand", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take spool of cat6", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take spool of cat6", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take spool of cat6", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take planetoid", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take sand", "take spool of cat6", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take sand", "take spool of cat6", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take sand", "take spool of cat6", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take sand", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take sand", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take sand", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take spool of cat6", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take spool of cat6", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take spool of cat6", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take fuel cell", "take coin", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take sand", "take spool of cat6", "take coin",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take sand", "take spool of cat6", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take sand", "take spool of cat6", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take sand", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take sand", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take sand", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take spool of cat6", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take spool of cat6", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take spool of cat6", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take planetoid", "take coin", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take sand", "take spool of cat6", "take coin", "take dark matter",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take sand", "take spool of cat6", "take coin", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take sand", "take spool of cat6", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take sand", "take coin", "take dark matter", "take wreath",
		"inv",
		"south",
		"drop jam", "drop fuel cell", "drop planetoid", "drop sand", "drop spool of cat6", "drop coin", "drop dark matter", "drop wreath",
		"take spool of cat6", "take coin", "take dark matter", "take wreath",
		"south",
	}

	// Convert commands to input stream
	var inputStream []int
	for _, cmd := range commands {
		for _, ch := range cmd + "\n" {
			inputStream = append(inputStream, int(ch))
		}
	}

	// Read program from file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	splitData := strings.Split(strings.TrimSpace(string(content)), ",")
	data := make([]int, len(splitData))
	for i, s := range splitData {
		data[i], _ = strconv.Atoi(s)
	}

	// Initialize and run the Intcode computer
	runner := NewIntcode(false)
	runner.initProgram(data, inputStream, true)
	runner.runProgram(nil)

	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for !runner.eop {
		for runner.readOutput {
			fmt.Print(string(rune(runner.retVal)))
			runner.runProgram(nil)
		}
		if runner.needsInput {
			fmt.Print("Input: ")
			input, _ := reader.ReadString('\n')
			inputChars := make([]int, len(input))
			for i, ch := range input {
				inputChars[i] = int(ch)
			}
			runner.runProgram(inputChars)
		}
	}
}
