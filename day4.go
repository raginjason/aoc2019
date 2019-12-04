package main

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

var adjacentDigits = [10]string{"00", "11", "22", "33", "44", "55", "66", "77", "88", "99"}

func CheckAdjacentDigitPassword(pass int) bool {
	passStr := strconv.Itoa(pass)

	for _, digits := range adjacentDigits {
		if strings.Contains(passStr, digits) {
			return true
		}
	}

	return false
}

func NonDecreasingDigitPassword(pass int) bool {
	passStr := strconv.Itoa(pass)

	for i := range passStr {
		if i+1 < len(passStr) && passStr[i] > passStr[i+1] {
			return false
		}
	}

	return true
}

func SixDigitPassword(pass int) bool {
	// Bottom of six digits
	if pass < 100000 {
		return false
	}

	// Top of six digits
	if pass > 999999 {
		return false
	}

	return true
}

// Clever or stupid? I'm not sure
var excessAdjacentDigits = [40]string{
	"000000", "111111", "222222", "333333", "444444", "555555", "666666", "777777", "888888", "999999",
	"00000", "11111", "22222", "33333", "44444", "55555", "66666", "77777", "88888", "99999",
	"0000", "1111", "2222", "3333", "4444", "5555", "6666", "7777", "8888", "9999",
	"000", "111", "222", "333", "444", "555", "666", "777", "888", "999",
}

func LargerGroupAdjacentDigitPassword(pass int) bool {
	passStr := strconv.Itoa(pass)

	for _, digits := range excessAdjacentDigits {
		// Removing digits could false positives, such as the case of 566655 -> 555
		// Because of this, replace with a space instead of "" just to be safe
		passStr = strings.Replace(passStr, digits, " ", -1)
	}

	for _, digits := range adjacentDigits {
		if strings.Contains(passStr, digits) {
			return true
		}
	}

	return false
}

func SecretContainer(begin, end int) int {
	passwords := make([]int, 0)
	for i := begin + 1; i < end; i = i + 1 {
		if SixDigitPassword(i) && CheckAdjacentDigitPassword(i) && NonDecreasingDigitPassword(i) {
			passwords = append(passwords, i)
		}
	}
	return len(passwords)
}

func BetterSecretContainer(begin, end int) int {
	passwords := make([]int, 0)
	for i := begin + 1; i < end; i = i + 1 {
		if SixDigitPassword(i) && CheckAdjacentDigitPassword(i) && NonDecreasingDigitPassword(i) && LargerGroupAdjacentDigitPassword(i) {
			passwords = append(passwords, i)
		}
	}
	return len(passwords)
}

func scanDay4File() []int {
	lines, err := FileInput(4)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	var passRange []int

	for _, line := range lines {
		r := csv.NewReader(strings.NewReader(line))
		r.Comma = '-'
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			for _, cell := range record {
				pr, err := strconv.Atoi(cell)
				if err != nil {
					log.Fatalf("could not convert %s, %s\n", cell, err)
				}
				passRange = append(passRange, pr)
			}
		}
	}
	return passRange
}

func day4() int {
	passRange := scanDay4File()

	res := SecretContainer(passRange[0], passRange[1])

	return res
}

func day4pt2() int {
	passRange := scanDay4File()

	res := BetterSecretContainer(passRange[0], passRange[1])

	return res
}
