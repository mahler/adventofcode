package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	type passportMapType map[string]string

	data, err := ioutil.ReadFile("passport.data")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	records := strings.Split(string(data), "\n\n")
	fmt.Println()
	fmt.Println("DAY04, Part 1: Passport Processing")
	fmt.Println("Total passwords in dataset: ", len(records))

	validpassport := 0

	passportDataset := make([]passportMapType, 0)

	for _, passportRecord := range records {
		//		fmt.Println(i, ":", passportRecord)
		passportRecord := strings.Split(passportRecord, "\n")

		//fmt.Println(i, ":", passportRecord)
		passportMap := make(map[string]string)
		for _, passportField := range passportRecord {

			//fmt.Println("singlerecord:", passportField, "- type:", reflect.TypeOf(passportField))
			if strings.Index(passportField, " ") < 0 {
				// Just a KeyValue pair on the row...
				tmp := strings.Split(passportField, ":")
				passportMap[tmp[0]] = string(tmp[1])
			} else {
				// Multiple records on row
				rowSplit := strings.Split(passportField, " ")
				for _, rowSplitData := range rowSplit {
					tmp := strings.Split(rowSplitData, ":")
					passportMap[tmp[0]] = string(tmp[1])
				}
			}

		}

		if _, ok := passportMap["byr"]; ok {
			if _, ok := passportMap["iyr"]; ok {
				if _, ok := passportMap["eyr"]; ok {
					if _, ok := passportMap["hgt"]; ok {
						if _, ok := passportMap["hcl"]; ok {
							if _, ok := passportMap["ecl"]; ok {
								if _, ok := passportMap["pid"]; ok {
									validpassport++
									passportDataset = append(passportDataset, passportMap)
								}
							}
						}
					}
				}
			}
		}

	}

	fmt.Println("Checked passports:", len(passportDataset))
	fmt.Println("Valid passports:", validpassport)

	fmt.Println()
	fmt.Println("DAY04, Part 2: Passport validation")
	validPassports := 0

	fmt.Println()
	for _, aPassport := range passportDataset {
		// Use passportCheckFail as string to embed debug info into it...
		passportCheckFail := ""

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		byr, err := strconv.Atoi(aPassport["byr"])
		if err != nil {
			passportCheckFail += "byr-missing - "
		} else if byr < 1920 || byr > 2002 {
			passportCheckFail += "byr-wrong - "
		}

		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		iyr, err := strconv.Atoi(aPassport["iyr"])
		if err != nil {
			passportCheckFail += "iyr-missing - "
		} else if iyr < 2010 || iyr > 2020 {
			passportCheckFail += "iyr-wrong - "
		}

		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		eyr, err := strconv.Atoi(aPassport["eyr"])
		if err != nil {
			passportCheckFail += "eyr-missing - "
		} else if eyr < 2020 || eyr > 2030 {
			passportCheckFail += "eyr-wrong - "
		}

		// hgt (Height) - a number followed by either cm or in:
		if aPassport["hgt"] != "" {
			unitHeight := aPassport["hgt"][len(aPassport["hgt"])-2:]
			// Remove unit from measurement string... and convert to number
			measurementStr := aPassport["hgt"][0 : len(aPassport["hgt"])-2]
			measurement, _ := strconv.Atoi(measurementStr)

			if unitHeight == "cm" {
				// If cm, the number must be at least 150 and at most 193.
				if measurement < 150 || measurement > 193 {
					passportCheckFail += "hgt-cm-wrong - "
				}

			} else if unitHeight == "in" {
				// If in, the number must be at least 59 and at most 76.
				if measurement < 59 || measurement > 76 {
					passportCheckFail += "hgt-in-wrong - "
				}

			} else {
				passportCheckFail += "hgt-missing - "
			}
		}

		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		if aPassport["ecl"] != "amb" && aPassport["ecl"] != "blu" && aPassport["ecl"] != "brn" && aPassport["ecl"] != "gry" && aPassport["ecl"] != "grn" && aPassport["ecl"] != "hzl" && aPassport["ecl"] != "oth" {
			passportCheckFail += "ecl-missing - "
		}

		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		if len(aPassport["pid"]) == 9 {
			// Should also check if number, but atlas...
			//fmt.Println(i, aPassport["pid"])
		} else {
			// not 9 digits
			passportCheckFail += "pid-missing - "
		}

		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		if aPassport["hcl"] != "" {
			hcl := aPassport["hcl"]

			if hcl[0:1] != "#" {
				passportCheckFail += "hcl-hash - "
			} else {
				color := []rune(hcl[1:])
				for _, cRune := range color {
					c := string(cRune)
					if c != "0" && c != "1" && c != "2" && c != "3" && c != "4" && c != "5" && c != "6" && c != "7" && c != "8" && c != "9" && c != "a" && c != "b" && c != "c" && c != "d" && c != "e" && c != "f" {
						passportCheckFail += "hcl-numberfail - "
					}
				}
			}
		}

		// No passport Check Failed if string is empty
		if passportCheckFail == "" {
			validPassports++
		} else {
			// Debug
			//fmt.Println(i, ":", passportCheckFail)
		}
	}

	fmt.Println("Validated Passports:", validPassports)
}
