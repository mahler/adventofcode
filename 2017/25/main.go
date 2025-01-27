package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TuringState represents the state name
type TuringState string

// Tape represents the machine's tape, mapping positions to boolean values
type Tape map[int]bool

// StateTransition represents what happens when a rule is triggered
type StateTransition struct {
	WriteValue   bool
	NewState     TuringState
	TapeMovement int
}

// RuleTrigger represents when a rule should be applied
type RuleTrigger struct {
	State TuringState
	Value bool
}

// Rules maps triggers to transitions
type Rules map[RuleTrigger]StateTransition

// Machine represents the complete state of the Turing machine
type Machine struct {
	State          TuringState
	Tape           Tape
	TapeLocation   int
	StepsRemaining int
}

// NewMachine creates an empty machine with initial settings
func NewMachine() Machine {
	return Machine{
		State:          "unknown",
		Tape:           make(Tape),
		TapeLocation:   0,
		StepsRemaining: 0,
	}
}

// ExecuteStep performs a single step of the machine
func (m *Machine) ExecuteStep(rules Rules) {
	// Get current tape value
	tapeValue, exists := m.Tape[m.TapeLocation]
	if !exists {
		tapeValue = false
	}

	// Find and apply transition
	trigger := RuleTrigger{m.State, tapeValue}
	transition, exists := rules[trigger]
	if !exists {
		panic(fmt.Sprintf("No rule found for state %s and value %v", m.State, tapeValue))
	}

	// Update tape
	m.Tape[m.TapeLocation] = transition.WriteValue

	// Update position
	m.TapeLocation += transition.TapeMovement

	// Update state
	m.State = transition.NewState

	// Decrement steps
	m.StepsRemaining--
}

// ExecuteAll runs the machine until no steps remain
func (m *Machine) ExecuteAll(rules Rules) {
	for m.StepsRemaining > 0 {
		m.ExecuteStep(rules)
	}
}

// ParseMachineDescription parses the input file format
func ParseMachineDescription(filename string) (Machine, Rules, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Machine{}, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(Rules)
	machine := NewMachine()

	// Parse initial state
	if !scanner.Scan() {
		return machine, rules, fmt.Errorf("unexpected end of file")
	}
	line := scanner.Text()
	parts := strings.Split(line, "state ")
	if len(parts) != 2 {
		return machine, rules, fmt.Errorf("invalid initial state line")
	}
	machine.State = TuringState(strings.Trim(parts[1], "."))

	// Parse number of steps
	if !scanner.Scan() {
		return machine, rules, fmt.Errorf("unexpected end of file")
	}
	line = scanner.Text()
	parts = strings.Split(line, "after ")
	if len(parts) != 2 {
		return machine, rules, fmt.Errorf("invalid steps line")
	}
	stepsStr := strings.Split(parts[1], " ")[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return machine, rules, err
	}
	machine.StepsRemaining = steps

	// Parse state transitions
	var currentState TuringState
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "In state") {
			parts = strings.Split(line, "state ")
			currentState = TuringState(strings.Trim(parts[1], ":"))
		} else if strings.HasPrefix(line, "If the current value is") {
			value := strings.Contains(line, "1")

			// Read write value
			scanner.Scan()
			writeValue := strings.Contains(scanner.Text(), "1")

			// Read movement
			scanner.Scan()
			movement := 1
			if strings.Contains(scanner.Text(), "left") {
				movement = -1
			}

			// Read next state
			scanner.Scan()
			parts = strings.Split(scanner.Text(), "state ")
			nextState := TuringState(strings.Trim(parts[1], ".-"))

			// Create transition
			trigger := RuleTrigger{currentState, value}
			transition := StateTransition{
				WriteValue:   writeValue,
				NewState:     nextState,
				TapeMovement: movement,
			}
			rules[trigger] = transition
		}
	}

	return machine, rules, nil
}

func main() {
	machine, rules, err := ParseMachineDescription("input.txt")
	if err != nil {
		fmt.Printf("Error parsing machine description: %v\n", err)
		return
	}

	machine.ExecuteAll(rules)

	// Count true values on tape
	count := 0
	for _, value := range machine.Tape {
		if value {
			count++
		}
	}
	fmt.Println("What is the diagnostic checksum it produces once it's working again?")
	fmt.Println(count)
}
