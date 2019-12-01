package main

import (
	"log"
	"strconv"
)

type computeFuel func(mass int) int

func ComputeFuelTotal(masses []int, fn computeFuel) int {
	sum := 0

	for i := range masses {
		sum += fn(masses[i])
	}

	return sum
}

func ComputeFuel(mass int) int {
	return int(mass/3) - 2
}

func ComputeRecursiveFuel(mass int) int {
	acc := ComputeFuel(mass)
	for v := ComputeFuel(acc); v > 0; v = ComputeFuel(v) {
		acc += v
	}
	return acc
}

func scanDay1File() []int {
	lines, err := FileInput(1)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	var numbers []int

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not convert %s, %s\n", line, err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func day1() int {
	masses := scanDay1File()

	return ComputeFuelTotal(masses, ComputeFuel)
}

func day1pt2() int {
	masses := scanDay1File()

	return ComputeFuelTotal(masses, ComputeRecursiveFuel)
}
