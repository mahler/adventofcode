package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// calculates the next secret.
func next(secret int64) int64 {
	secret = ((secret * 64) ^ secret) % 16777216
	secret = ((secret / 32) ^ secret) % 16777216
	return ((secret * 2048) ^ secret) % 16777216
}

// calculates the nth secret value.
func nthSecret(secret, n int64) int64 {
	for i := int64(0); i < n; i++ {
		secret = next(secret)
	}
	return secret
}

// calculates sequences of price differences.
func nSequences(secret, n int64) map[string]int64 {
	prices := []int64{secret % 10}
	state := secret
	for i := int64(0); i < n; i++ {
		state = next(state)
		prices = append(prices, state%10)
	}

	differences := []int64{}
	for i := 0; i < len(prices)-1; i++ {
		differences = append(differences, prices[i+1]-prices[i])
	}

	sequences := map[string]int64{}
	for i := 0; i < len(differences)-3; i++ {
		key := fmt.Sprintf("%v,%v,%v,%v", differences[i], differences[i+1], differences[i+2], differences[i+3])
		if _, exists := sequences[key]; !exists {
			sequences[key] = prices[i+4]
		}
	}

	return sequences
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	secrets := []int64{}
	for _, line := range lines {
		secret, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			secrets = append(secrets, secret)
		}
	}

	// Part 1
	p1 := int64(0)
	for _, secret := range secrets {
		p1 += nthSecret(secret, 2000)
	}
	fmt.Println("Part 1: What is the sum of the 2000th secret number generated by each buyer?")
	fmt.Println(p1)

	// Part 2
	var mu sync.Mutex
	counts := map[string]int64{}

	var wg sync.WaitGroup
	for _, secret := range secrets {
		wg.Add(1)
		go func(secret int64) {
			defer wg.Done()
			seqs := nSequences(secret, 2000)
			mu.Lock()
			for k, v := range seqs {
				counts[k] += v
			}
			mu.Unlock()
		}(secret)
	}
	wg.Wait()

	maxCount := int64(0)
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
	}

	if len(counts) == 0 {
		fmt.Println("Error: no max value found")
		return
	}

	fmt.Println()
	fmt.Println("Part 2: What is the most bananas you can get?")
	fmt.Println(maxCount)
}