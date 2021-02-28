package main

// Most code here sourced from https://github.com/devries/advent_of_code_2020/blob/main/day18_p1/main.go
// and https://github.com/devries/advent_of_code_2020/blob/main/day18_p2/main.go

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type TokenType int

const (
	NUMBER TokenType = iota
	MULOPER
	ADDOPER
	LPARENS
	RPARENS
)

type Token struct {
	token TokenType
	value int
}

type Statement struct {
	tokens   []Token
	position int
}

func main() {
	var sum int64
	data, err := os.ReadFIle("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)

	}
	strSlice := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, row := range strSlice {
		stmt := tokenize(row)
		sum += int64(evalExpression(&stmt))
	}

	fmt.Println("Day 18 - Part 1: Day 18: Operation Order")
	fmt.Println("Sum of calculus across lines:", sum)

	fmt.Println()
	fmt.Println("Part 2")
	var p2sum int64
	for _, row := range strSlice {
		stmt := tokenize(row)
		p2sum += int64(p2evalExpression(&stmt))
	}
	fmt.Println(p2sum)
}

func tokenize(input string) Statement {
	toks := []Token{}
	var b strings.Builder

	for _, c := range input {
		if c >= '0' && c <= '9' {
			// Number
			b.WriteRune(c)
		} else {
			if b.Len() > 0 {
				v, err := strconv.Atoi(b.String())
				if err != nil {
					panic(fmt.Errorf("Unable to parse numbers in %s", input))
				}
				toks = append(toks, Token{NUMBER, v})
				b.Reset()
			}
			switch c {
			case '+':
				toks = append(toks, Token{ADDOPER, 0})
			case '*':
				toks = append(toks, Token{MULOPER, 0})
			case '(':
				toks = append(toks, Token{LPARENS, 0})
			case ')':
				toks = append(toks, Token{RPARENS, 0})
			}
		}
	}
	if b.Len() > 0 {
		v, err := strconv.Atoi(b.String())
		if err != nil {
			panic(fmt.Errorf("Unable to parse numbers in %s", input))
		}
		toks = append(toks, Token{NUMBER, v})
	}

	return Statement{toks, 0}
}

func evalExpression(stmt *Statement) int {
	left := evalTerm(stmt)

	for stmt.position < len(stmt.tokens) {
		current := stmt.tokens[stmt.position]
		switch current.token {
		case ADDOPER:
			stmt.position++
			right := evalTerm(stmt)
			left = left + right
		case MULOPER:
			stmt.position++
			right := evalTerm(stmt)
			left = left * right
		default:
			return left
		}
	}

	return left
}

func evalTerm(stmt *Statement) int {
	current := stmt.tokens[stmt.position]

	if current.token == NUMBER {
		stmt.position++
		return current.value
	}

	if current.token == LPARENS {
		stmt.position++
		v := evalExpression(stmt)
		current = stmt.tokens[stmt.position]
		if current.token != RPARENS {
			panic(fmt.Errorf("Unbalanced parenthesis"))
		}
		stmt.position++
		return v
	}

	panic(fmt.Errorf("No term found"))
}

func p2evalExpression(stmt *Statement) int {
	left := p2evalFactor(stmt)

	for stmt.position < len(stmt.tokens) {
		current := stmt.tokens[stmt.position]
		switch current.token {
		case MULOPER:
			stmt.position++
			right := p2evalFactor(stmt)
			left = left * right
		default:
			return left
		}
	}

	return left
}

func p2evalFactor(stmt *Statement) int {
	left := p2evalTerm(stmt)

	for stmt.position < len(stmt.tokens) {
		current := stmt.tokens[stmt.position]
		switch current.token {
		case ADDOPER:
			stmt.position++
			right := p2evalTerm(stmt)
			left = left + right
		default:
			return left
		}
	}

	return left
}

func p2evalTerm(stmt *Statement) int {
	current := stmt.tokens[stmt.position]

	if current.token == NUMBER {
		stmt.position++
		return current.value
	}

	if current.token == LPARENS {
		stmt.position++
		v := p2evalExpression(stmt)
		current = stmt.tokens[stmt.position]
		if current.token != RPARENS {
			panic(fmt.Errorf("Unbalanced parenthesis"))
		}
		stmt.position++
		return v
	}

	panic(fmt.Errorf("No term found"))
}
