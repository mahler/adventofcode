package main

import (
	"testing"
)

func TestMassToFuel(t *testing.T) {
	m2fTestData := make(map[int]int)
	m2fTestData[12] = 2
	m2fTestData[14] = 2
	m2fTestData[1969] = 654
	m2fTestData[100756] = 33583

	for mass, fuelNeed := range m2fTestData {
		result := massToFuel(mass)
		if result != fuelNeed {
			t.Errorf("Wrong result for %v - got %v, wanted %v", mass, result, fuelNeed)
		}

	}
}

func TestMassToFuelWithFuel(t *testing.T) {
	m2fTestData := make(map[int]int)
	m2fTestData[14] = 2
	m2fTestData[1969] = 966
	m2fTestData[100756] = 50346

	for mass, fuelNeed := range m2fTestData {
		result := massToFuelWithFuel(mass)
		if result != fuelNeed {
			t.Errorf("Wrong result for %v - got %v, wanted %v", mass, result, fuelNeed)
		}
	}
}
