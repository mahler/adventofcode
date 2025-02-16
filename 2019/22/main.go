package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// primeModInv calculates the modular multiplicative inverse for prime p
func primeModInv(p, mod *big.Int) *big.Int {
	result := new(big.Int)
	return result.Exp(p, result.Sub(mod, big.NewInt(2)), mod)
}

// function represents a linear function ax + b
type function struct {
	a, b *big.Int
}

// inverseFunctions generates inverse mappings for each shuffle instruction
func inverseFunctions(instructions []string, ncards *big.Int) []function {
	result := make([]function, 0, len(instructions))

	for i := len(instructions) - 1; i >= 0; i-- {
		line := instructions[i]
		if strings.HasPrefix(line, "deal into new stack") {
			// f(x) = -x + (ncards-1)
			result = append(result, function{
				a: big.NewInt(-1),
				b: new(big.Int).Sub(ncards, big.NewInt(1)),
			})
		} else if strings.HasPrefix(line, "cut") {
			// f(x) = x + amount
			amount, _ := strconv.ParseInt(strings.Fields(line)[1], 10, 64)
			result = append(result, function{
				a: big.NewInt(1),
				b: big.NewInt(amount),
			})
		} else {
			// f(x) = increment^(-1) * x + 0
			increment, _ := strconv.ParseInt(strings.Fields(line)[len(strings.Fields(line))-1], 10, 64)
			inc := big.NewInt(increment)
			result = append(result, function{
				a: primeModInv(inc, ncards),
				b: big.NewInt(0),
			})
		}
	}
	return result
}

// fcompose composes two linear functions mod m
func fcompose(f, g function, mod *big.Int) function {
	a := new(big.Int).Mul(g.a, f.a)
	a.Mod(a, mod)

	b := new(big.Int).Mul(g.a, f.b)
	b.Add(b, g.b)
	b.Mod(b, mod)

	return function{a: a, b: b}
}

// inverseFunction composes all inverse mappings into a single function
func inverseFunction(instructions []string, ncards *big.Int) function {
	funcs := inverseFunctions(instructions, ncards)
	result := function{a: big.NewInt(1), b: big.NewInt(0)}

	for _, f := range funcs {
		result = fcompose(result, f, ncards)
	}
	return result
}

// fapply computes f(x) = ax + b
func fapply(f function, x, mod *big.Int) *big.Int {
	result := new(big.Int).Mul(f.a, x)
	result.Add(result, f.b)
	return result.Mod(result, mod)
}

// frepeat repeats function f n times
func frepeat(f function, n, mod *big.Int) function {
	if n.Cmp(big.NewInt(0)) == 0 {
		return function{a: big.NewInt(1), b: big.NewInt(0)}
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return f
	}

	half := new(big.Int).Div(n, big.NewInt(2))
	odd := new(big.Int).Mod(n, big.NewInt(2))

	g := frepeat(f, half, mod)
	gg := fcompose(g, g, mod)

	if odd.Cmp(big.NewInt(1)) == 0 {
		return fcompose(f, gg, mod)
	}
	return gg
}

// findCard finds the position that will end up at the given value
func findCard(value, ncards *big.Int, instructions []string) *big.Int {
	f := inverseFunction(instructions, ncards)

	for i := new(big.Int); i.Cmp(ncards) < 0; i.Add(i, big.NewInt(1)) {
		if fapply(f, i, ncards).Cmp(value) == 0 {
			return i
		}
	}
	return nil
}

// cardAt finds the value at a given position after n shuffles
func cardAt(index, ncards, nshuffles *big.Int, instructions []string) *big.Int {
	f := inverseFunction(instructions, ncards)
	fn := frepeat(f, nshuffles, ncards)
	return fapply(fn, index, ncards)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	// First part
	ncards := big.NewInt(10007)
	value := big.NewInt(2019)
	result1 := findCard(value, ncards, instructions)
	fmt.Println("Part 1: After shuffling your factory order deck of 10007 cards, what is the position of card 2019?")
	fmt.Println(result1)

	// Second part
	ncards2 := new(big.Int)
	ncards2.SetString("119315717514047", 10)
	nshuffles := new(big.Int)
	nshuffles.SetString("101741582076661", 10)
	index := big.NewInt(2020)
	result2 := cardAt(index, ncards2, nshuffles, instructions)
	fmt.Println()
	fmt.Println("Part 2: After shuffling your new, giant, factory order deck that many times, what number is on the card that ends up in position 2020?")
	fmt.Println(result2)
}
