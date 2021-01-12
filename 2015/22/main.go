package main

import (
	"fmt"
)

type spell struct {
	cost     int
	hitpoint int
	mana     int
	damage   int
	armor    int
	duration int
}

type player struct {
	hitpoint int
	damage   int
	armor    int
	mana     int
	spells   map[string]spell
}

func main() {
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 22, Part 1: Wizard Simulator 20XX")
	part2 := false
	min := 10000
	for i := 1; min == 10000; i++ {
		spells := map[string]spell{
			"Magic Missile": spell{53, 0, 0, 4, 0, -1},
			"Drain":         spell{73, 2, 0, 2, 0, -1},
			"Shield":        spell{113, 0, 0, 0, 7, 6},
			"Poison":        spell{173, 0, 0, 3, 0, 6},
			"Recharge":      spell{229, 0, 101, 0, 0, 5},
		}

		p1 := player{
			hitpoint: 50,
			mana:     500,
		}

		boss := player{
			hitpoint: 51,
			damage:   9,
		}

		min = play(0, p1, boss, spells, min, i, part2)
	}
	fmt.Println("What is the least amount of mana you can spend and still win the fight?")
	fmt.Println(min)
	// Part 2 -----------------
	fmt.Println()
	fmt.Println("Part 2")
	part2 = true
	min = 10000
	for i := 1; min == 10000; i++ {
		spells := map[string]spell{
			"Magic Missile": spell{53, 0, 0, 4, 0, -1},
			"Drain":         spell{73, 2, 0, 2, 0, -1},
			"Shield":        spell{113, 0, 0, 0, 7, 6},
			"Poison":        spell{173, 0, 0, 3, 0, 6},
			"Recharge":      spell{229, 0, 101, 0, 0, 5},
		}

		p1 := player{
			hitpoint: 50,
			mana:     500,
		}

		boss := player{
			hitpoint: 51,
			damage:   9,
		}

		min = play(0, p1, boss, spells, min, i, part2)
	}

	fmt.Println("With the same starting stats for you and the boss,")
	fmt.Println("what is the least amount of mana you can spend and still win the fight?")
	fmt.Println(min)
}

func play(mana int, you, boss player, spells map[string]spell, max, depth int, part2 bool) int {
	result := 10000
	if depth == 0 || mana > max {
		return result
	}
	for name, spell := range spells {
		nyou, nboss, ended := cast(name, spell, you, boss, part2)
		if ended {
			if nboss.hitpoint <= 0 {
				if result > mana+spell.cost {
					result = mana + spell.cost
				}
			}
			continue
		}
		nboss, nyou, ended = attack(nboss, nyou)
		if ended {
			if nboss.hitpoint <= 0 {
				if result > mana+spell.cost {
					result = mana + spell.cost
				}
			}
			continue
		}

		playGame := play(mana+spell.cost, nyou, nboss, spells, max, depth-1, part2)
		if playGame < result {
			result = playGame
		}
	}
	return result
}

func attack(player, defender player) (player, player, bool) {
	defender, player = timer(defender, player)
	if player.hitpoint <= 0 {
		return player, defender, true
	}
	damage := player.damage - defender.armor
	if damage < 1 {
		damage = 1
	}
	defender.hitpoint -= damage
	return player, defender, defender.hitpoint <= 0
}

func cast(name string, spell spell, player, defender player, part2 bool) (player, player, bool) {
	player, defender = timer(player, defender)
	// ---- P2
	if part2 {
		player.hitpoint--
		if player.hitpoint <= 0 {
			return player, defender, true
		}
	}
	//----
	if defender.hitpoint <= 0 {
		return player, defender, true
	}
	if _, exists := player.spells[name]; exists {
		return player, defender, true
	}
	if player.mana < spell.cost {
		return player, defender, true
	}
	player.mana -= spell.cost
	if spell.duration == -1 {
		player.hitpoint += spell.hitpoint
		defender.hitpoint -= spell.damage
	} else {
		player.spells[name] = spell
	}
	return player, defender, defender.hitpoint <= 0
}

func timer(player, defender player) (player, player) {
	spells := map[string]spell{}
	player.armor = 0
	for name, spell := range player.spells {
		player.hitpoint += spell.hitpoint
		player.mana += spell.mana
		defender.hitpoint -= spell.damage
		player.armor += spell.armor
		spell.duration--
		if spell.duration > 0 {
			spells[name] = spell
		}
	}
	player.spells = spells
	return player, defender
}
