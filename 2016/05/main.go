package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	input := "reyedfim"

	fmt.Println()
	fmt.Println("2016")
	fmt.Println("Day 05, Part 1: How About a Nice Game of Chess?")

	password := ""
	suffix := 0
	for pwIndex := 0; pwIndex < 8; suffix++ {

		testString := fmt.Sprintf("%s%d", input, suffix)
		hash := md5.New()
		io.WriteString(hash, testString)
		hashString := fmt.Sprintf("%x", hash.Sum(nil))

		if strings.Index(hashString, "00000") == 0 {
			password += string(hashString[5])
			pwIndex++
		}
	}
	fmt.Println("Cracked password:")
	fmt.Println(password)

	// ----------------------------------------------------------------
	fmt.Println()
	fmt.Println("Part 2: Note - this is quite slow. Please be patient...")

	var newPassword [8]rune
	var passwordCount int

	for suffix := 0; passwordCount < 8; suffix++ {
		testString := fmt.Sprintf("%s%d", input, suffix)
		hash := md5.New()
		io.WriteString(hash, testString)
		hashString := fmt.Sprintf("%x", hash.Sum(nil))

		if strings.Index(hashString, "00000") == 0 {
			index, err := strconv.Atoi(hashString[5:6])
			if err != nil || index >= 8 || newPassword[index] != rune(0) {
				continue
			}

			passwordRune := rune(hashString[6])
			newPassword[index] = passwordRune
			passwordCount++
		}
	}

	fmt.Println(string(newPassword[:]))
}
