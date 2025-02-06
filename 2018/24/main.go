package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Unit struct {
	id     string
	n      int
	hp     int
	immune map[string]bool
	weak   map[string]bool
	init   int
	dtyp   string
	dmg    int
	side   int
	target *Unit
}

func (u *Unit) power() int {
	return u.n * u.dmg
}

func (u *Unit) dmgTo(v *Unit) int {
	if v.immune[u.dtyp] {
		return 0
	} else if v.weak[u.dtyp] {
		return 2 * u.power()
	}
	return u.power()
}

func processResistances(s string) (map[string]bool, map[string]bool) {
	immune := make(map[string]bool)
	weak := make(map[string]bool)

	parts := strings.Split(s, " ")
	resistType := parts[0]
	target := immune
	if resistType == "weak" {
		target = weak
	}

	for _, word := range parts[2:] {
		word = strings.TrimSuffix(word, ",")
		target[word] = true
	}

	return immune, weak
}

func battle(originalUnits []*Unit, boost int) (int, int) {
	units := make([]*Unit, len(originalUnits))
	for i, u := range originalUnits {
		newDmg := u.dmg
		if u.side == 0 {
			newDmg += boost
		}
		units[i] = &Unit{
			id:     u.id,
			n:      u.n,
			hp:     u.hp,
			immune: u.immune,
			weak:   u.weak,
			init:   u.init,
			dtyp:   u.dtyp,
			dmg:    newDmg,
			side:   u.side,
		}
	}

	for {
		// Sort units by power and initiative
		sort.Slice(units, func(i, j int) bool {
			if units[i].power() != units[j].power() {
				return units[i].power() > units[j].power()
			}
			return units[i].init > units[j].init
		})

		chosen := make(map[string]bool)
		for _, u := range units {
			var targets []*Unit
			for _, v := range units {
				if v.side != u.side && !chosen[v.id] && u.dmgTo(v) > 0 {
					targets = append(targets, v)
				}
			}

			if len(targets) > 0 {
				sort.Slice(targets, func(i, j int) bool {
					dmgI := u.dmgTo(targets[i])
					dmgJ := u.dmgTo(targets[j])
					if dmgI != dmgJ {
						return dmgI > dmgJ
					}
					if targets[i].power() != targets[j].power() {
						return targets[i].power() > targets[j].power()
					}
					return targets[i].init > targets[j].init
				})
				u.target = targets[0]
				chosen[targets[0].id] = true
			}
		}

		// Sort by initiative for attack phase
		sort.Slice(units, func(i, j int) bool {
			return units[i].init > units[j].init
		})

		anyKilled := false
		for _, u := range units {
			if u.target != nil {
				dmg := u.dmgTo(u.target)
				killed := min(u.target.n, dmg/u.target.hp)
				if killed > 0 {
					anyKilled = true
				}
				u.target.n -= killed
			}
		}

		// Remove dead units and clear targets
		var newUnits []*Unit
		for _, u := range units {
			if u.n > 0 {
				u.target = nil
				newUnits = append(newUnits, u)
			}
		}
		units = newUnits

		if !anyKilled {
			n1 := 0
			for _, u := range units {
				if u.side == 1 {
					n1 += u.n
				}
			}
			return 1, n1
		}

		n0, n1 := 0, 0
		for _, u := range units {
			if u.side == 0 {
				n0 += u.n
			} else {
				n1 += u.n
			}
		}

		if n0 == 0 {
			return 1, n1
		}
		if n1 == 0 {
			return 0, n0
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var units []*Unit
	scanner := bufio.NewScanner(file)
	nextID := 1
	side := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Immune System") {
			nextID = 1
			side = 0
			continue
		}
		if strings.Contains(line, "Infection") {
			nextID = 1
			side = 1
			continue
		}
		if line == "" {
			continue
		}

		words := strings.Fields(line)
		n, _ := strconv.Atoi(words[0])
		hp, _ := strconv.Atoi(words[4])

		immune := make(map[string]bool)
		weak := make(map[string]bool)

		if strings.Contains(line, "(") {
			parts := strings.Split(line, "(")
			resists := strings.Split(parts[1], ")")[0]
			if strings.Contains(resists, ";") {
				resistParts := strings.Split(resists, ";")
				i1, w1 := processResistances(strings.TrimSpace(resistParts[0]))
				i2, w2 := processResistances(strings.TrimSpace(resistParts[1]))
				for k, v := range i1 {
					immune[k] = v
				}
				for k, v := range i2 {
					immune[k] = v
				}
				for k, v := range w1 {
					weak[k] = v
				}
				for k, v := range w2 {
					weak[k] = v
				}
			} else {
				immune, weak = processResistances(resists)
			}
		}

		init, _ := strconv.Atoi(words[len(words)-1])
		dtyp := words[len(words)-5]
		dmg, _ := strconv.Atoi(words[len(words)-6])

		name := fmt.Sprintf("%s_%d", map[int]string{1: "Infection", 0: "System"}[side], nextID)
		units = append(units, &Unit{
			id:     name,
			n:      n,
			hp:     hp,
			immune: immune,
			weak:   weak,
			init:   init,
			dtyp:   dtyp,
			dmg:    dmg,
			side:   side,
		})
		nextID++
	}

	// Part 1
	_, remaining := battle(units, 0)
	fmt.Println("Part 1: As it stands now, how many units would the winning army have?")
	fmt.Println(remaining)

	// Part 2
	boost := 0
	for {
		winner, remaining := battle(units, boost)
		if winner == 0 {
			fmt.Println()
			fmt.Println("Part 2: How many units does the immune system have left after getting the smallest boost it needs to win?")
			fmt.Println(remaining)
			break
		}
		boost++
	}
}
