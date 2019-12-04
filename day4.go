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

func SecretContainer(begin, end int) int {
	passwords := make([]int, 0)
	for i := begin + 1; i < end; i = i + 1 {
		if SixDigitPassword(i) && CheckAdjacentDigitPassword(i) && NonDecreasingDigitPassword(i) {
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
