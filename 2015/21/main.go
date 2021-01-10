package main

import (
	"fmt"
)

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type player struct {
	hitpoint   int
	damage     int
	armor      int
	configCost int
}

func main() {

	// Item Grid
	weapons := []item{
		item{name: "Dagger", cost: 8, damage: 4, armor: 0},
		item{name: "Shortsword", cost: 10, damage: 5, armor: 0},
		item{name: "Warhammer", cost: 25, damage: 6, armor: 0},
		item{name: "Longsword", cost: 40, damage: 7, armor: 0},
		item{name: "Greataxe", cost: 74, damage: 8, armor: 0},
	}
	armor := []item{
		item{name: "Leather", cost: 13, damage: 0, armor: 1},
		item{name: "Chainmail", cost: 31, damage: 0, armor: 2},
		item{name: "Splintmail", cost: 53, damage: 0, armor: 3},
		item{name: "Bandedmail", cost: 75, damage: 0, armor: 4},
		item{name: "Platemail", cost: 102, damage: 0, armor: 5},
		// armour is optional make fake item
		item{name: "none", cost: 0, damage: 0, armor: 0},
	}

	rings := []item{
		item{name: "Damage +1", cost: 25, damage: 1, armor: 0},
		item{name: "Damage +2", cost: 50, damage: 2, armor: 0},
		item{name: "Damage +3", cost: 100, damage: 3, armor: 0},
		item{name: "Defence +1", cost: 20, damage: 0, armor: 1},
		item{name: "Defence +2", cost: 40, damage: 0, armor: 2},
		item{name: "Defence +3", cost: 80, damage: 0, armor: 3},
		// rings are optional
		item{name: "None", cost: 0, damage: 0, armor: 0},
	}

	// Part 1 --------------
	fmt.Println()
	fmt.Println("2015")
	fmt.Println("Day 21, part 1: RPG Simulator 20XX")

	// Weapon 1 * Armor 0|1 * Rings 0-2
	config := [][]item{}

	for _, w := range weapons {
		for _, a := range armor {
			for _, r1 := range rings {
				for _, r2 := range rings {
					newConfig := []item{}
					newConfig = append(newConfig, w)
					newConfig = append(newConfig, a)
					newConfig = append(newConfig, r1)
					if r1 != r2 {
						newConfig = append(newConfig, r2)
					}
					config = append(config, newConfig)
				}
			}

		}
	}
	//	fmt.Println(len(config))
	boss := player{hitpoint: 100, damage: 8, armor: 2}

	// leastCost init with some high random number, as the assumed result cost is well below.
	leastCost := 99999
	maxCost := 0

	// Let's play every config setup...
	for _, playConfig := range config {
		player := player{hitpoint: 100}
		for _, configItem := range playConfig {
			player.armor += configItem.armor
			player.damage += configItem.damage
			player.configCost += configItem.cost
		}

		if playerWinsGame(player, boss) {
			if player.configCost < leastCost {
				leastCost = player.configCost
			}
		} else { // Part 2
			if player.configCost > maxCost {
				maxCost = player.configCost
			}
		}

		//fmt.Println("config: armour/", player.armor, "damage/", player.damage, "cost/", player.configCost)
		//break

	}

	fmt.Println("What is the least amount of gold you can spend and still win the fight?")
	fmt.Println(leastCost)
	// -------------------------
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("What is the most amount of gold you can spend and still lose the fight?")
	fmt.Println(maxCost)

}

func playerWinsGame(p1 player, boss player) bool {
	playerDamage := p1.damage - boss.armor
	if playerDamage < 1 {
		playerDamage = 1
	}
	bossDamage := boss.damage - p1.armor
	if bossDamage < 1 {
		bossDamage = 1
	}

	for {
		// Player attack boss
		boss.hitpoint -= playerDamage
		if boss.hitpoint <= 0 {
			break
		}

		// Boss attack player
		p1.hitpoint -= bossDamage

		if p1.hitpoint <= 0 {
			break
		}
	}

	return p1.hitpoint > 0
}
