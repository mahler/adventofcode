package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	FREE   = "."
	ELF    = "E"
	GOBLIN = "G"

	DEFAULT_ATTACK = 3
	DEFAULT_HP     = 200
)

type Position struct {
	X int
	Y int
}

type Player struct {
	Type   string
	HP     int
	Attack int
	Alive  bool
	Pos    Position
}

type GameData struct {
	Map       [][]string
	Players   []Player
	Round     int
	Outcome   int
	ElfAttack int
}

func main() {
	MAP := readInput("input.txt") // Replace with your input file

	// Part 1
	result := runBattle(deepCopyMap(MAP), DEFAULT_ATTACK, false)
	fmt.Println("Part 1: What is the outcome of the combat described in your puzzle input?")
	fmt.Println(result.Outcome)

	// Part 2
	var part2result *GameData
	elfAttack := DEFAULT_ATTACK
	for {
		elfAttack++
		part2result = runBattle(deepCopyMap(MAP), elfAttack, true)
		if part2result != nil {
			break
		}
	}
	fmt.Println()
	fmt.Println("Part 2: After increasing the Elves' attack power until it is just barely enough for")
	fmt.Println("them to win without any Elves dying, what is the outcome of the combat described in your puzzle input?")
	fmt.Println(part2result.Outcome)
}

func readInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var MAP [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		MAP = append(MAP, row)
	}
	return MAP
}

func runBattle(m [][]string, elfAttack int, abortIfElfDies bool) *GameData {
	data := &GameData{
		Map:       m,
		Players:   initPlayers(m, elfAttack),
		Round:     0,
		Outcome:   0,
		ElfAttack: elfAttack,
	}

	for {
		// Sort players by reading order
		sort.Slice(data.Players, func(i, j int) bool {
			if data.Players[i].Pos.Y == data.Players[j].Pos.Y {
				return data.Players[i].Pos.X < data.Players[j].Pos.X
			}
			return data.Players[i].Pos.Y < data.Players[j].Pos.Y
		})

		for i := range data.Players {
			if !data.Players[i].Alive {
				continue
			}

			// Check if there are any enemies left
			hasEnemies := false
			for _, p := range data.Players {
				if p.Alive && p.Type != data.Players[i].Type {
					hasEnemies = true
					break
				}
			}
			if !hasEnemies {
				// Calculate outcome
				sum := 0
				for _, p := range data.Players {
					if p.Alive {
						sum += p.HP
					}
				}
				data.Outcome = data.Round * sum
				return data
			}

			enemy := findEnemyToAttack(&data.Players[i], data.Players)
			var next *Position
			if enemy == nil {
				next = findNextMovement(&data.Players[i], data.Players, data.Map)
			}

			if enemy == nil && next != nil {
				// Move
				data.Map[data.Players[i].Pos.Y][data.Players[i].Pos.X] = FREE
				data.Players[i].Pos = *next
				data.Map[data.Players[i].Pos.Y][data.Players[i].Pos.X] = data.Players[i].Type

				// Check for enemy again after moving
				enemy = findEnemyToAttack(&data.Players[i], data.Players)
			}

			if enemy != nil {
				// Attack
				enemy.HP -= data.Players[i].Attack
				if enemy.HP < 1 {
					enemy.Alive = false
					data.Map[enemy.Pos.Y][enemy.Pos.X] = FREE
					if enemy.Type == ELF && abortIfElfDies {
						return nil
					}
				}
			}
		}
		data.Round++
	}
}

func initPlayers(m [][]string, elfAttack int) []Player {
	var players []Player
	for y := range m {
		for x := range m[y] {
			if m[y][x] == ELF || m[y][x] == GOBLIN {
				attack := DEFAULT_ATTACK
				if m[y][x] == ELF {
					attack = elfAttack
				}
				players = append(players, Player{
					Type:   m[y][x],
					HP:     DEFAULT_HP,
					Attack: attack,
					Alive:  true,
					Pos:    Position{X: x, Y: y},
				})
			}
		}
	}
	return players
}

func findEnemyToAttack(player *Player, allPlayers []Player) *Player {
	var weakest *Player
	minHP := DEFAULT_HP + 1

	for i := range allPlayers {
		p := &allPlayers[i]
		if !p.Alive || p.Type == player.Type {
			continue
		}

		// Check if in range
		inRange := (abs(p.Pos.X-player.Pos.X) == 1 && p.Pos.Y == player.Pos.Y) ||
			(abs(p.Pos.Y-player.Pos.Y) == 1 && p.Pos.X == player.Pos.X)

		if inRange && (weakest == nil || p.HP < minHP) {
			weakest = p
			minHP = p.HP
		}
	}
	return weakest
}

func findNextMovement(player *Player, allPlayers []Player, m [][]string) *Position {
	targetMap := make(map[string]Position)

	// Find all possible target positions
	for _, p := range allPlayers {
		if !p.Alive || p.Type == player.Type {
			continue
		}
		for _, adj := range getAdjacents(p.Pos) {
			if adj.Y >= 0 && adj.Y < len(m) && adj.X >= 0 && adj.X < len(m[adj.Y]) {
				if m[adj.Y][adj.X] == FREE {
					key := fmt.Sprintf("%d,%d", adj.X, adj.Y)
					targetMap[key] = adj
				}
			}
		}
	}

	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d", player.Pos.X, player.Pos.Y)] = true

	type Path []Position
	paths := []Path{{player.Pos}}

	for len(paths) > 0 {
		var newPaths []Path
		var targetPaths []Path

		for _, path := range paths {
			last := path[len(path)-1]
			for _, adj := range getAdjacents(last) {
				if adj.Y < 0 || adj.Y >= len(m) || adj.X < 0 || adj.X >= len(m[adj.Y]) {
					continue
				}

				key := fmt.Sprintf("%d,%d", adj.X, adj.Y)
				if _, exists := targetMap[key]; exists {
					newPath := append(Path(nil), path...)
					newPath = append(newPath, adj, targetMap[key])
					targetPaths = append(targetPaths, newPath)
				} else if !visited[key] && m[adj.Y][adj.X] == FREE {
					newPath := append(Path(nil), path...)
					newPath = append(newPath, adj)
					newPaths = append(newPaths, newPath)
				}
				visited[key] = true
			}
		}

		if len(targetPaths) > 0 {
			// Sort target paths by reading order of final position
			sort.Slice(targetPaths, func(i, j int) bool {
				pi := targetPaths[i][len(targetPaths[i])-1]
				pj := targetPaths[j][len(targetPaths[j])-1]
				if pi.Y == pj.Y {
					return pi.X < pj.X
				}
				return pi.Y < pj.Y
			})
			return &targetPaths[0][1]
		}

		paths = newPaths
	}

	return nil
}

func getAdjacents(pos Position) []Position {
	return []Position{
		{X: pos.X, Y: pos.Y - 1},
		{X: pos.X - 1, Y: pos.Y},
		{X: pos.X + 1, Y: pos.Y},
		{X: pos.X, Y: pos.Y + 1},
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func deepCopyMap(m [][]string) [][]string {
	newMap := make([][]string, len(m))
	for i := range m {
		newMap[i] = make([]string, len(m[i]))
		copy(newMap[i], m[i])
	}
	return newMap
}
