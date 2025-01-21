package main

import (
	"fmt"
)

type Gen struct {
	prev   int
	factor int
	limit  int
}

func NewGen(input, factor, limit int) *Gen {
	return &Gen{
		prev:   input,
		factor: factor,
		limit:  limit,
	}
}

func (g *Gen) Next() int {
	g.prev = g.prev * g.factor % 2147483647
	return g.prev
}

func (g *Gen) NextWithLimit() int {
	for {
		val := g.Next()
		if val%g.limit == 0 {
			return val
		}
	}
}

func main() {
	// Part 1
	a := NewGen(591, 16807, 4)
	b := NewGen(393, 48271, 8)
	c := 0

	for i := 0; i < 40000000; i++ {
		av := a.Next() & 0xFFFF
		bv := b.Next() & 0xFFFF

		if av == bv {
			c++
		}
	}

	fmt.Println("Part 1: After 40 million pairs, what is the judge's final count?")
	fmt.Println(c)

	// Part 2
	a = NewGen(591, 16807, 4)
	b = NewGen(393, 48271, 8)
	c = 0

	for i := 0; i < 5000000; i++ {
		av := a.NextWithLimit() & 0xFFFF
		bv := b.NextWithLimit() & 0xFFFF

		if av == bv {
			c++
		}
	}

	fmt.Println()
	fmt.Println("Part 2: After 5 million pairs, but using this new generator logic, what is the judge's final count?")
	fmt.Println(c)
}
