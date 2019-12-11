package main

import (
	"fmt"
	"github.com/raginjason/aoc2019/day5"
)

func main() {
	day1 := day1()
	fmt.Printf("Day 1 answer: %d\n", day1)
	day1pt2 := day1pt2()
	fmt.Printf("Day 1 part 2 answer: %d\n", day1pt2)

	day2 := day2()
	fmt.Printf("Day 2 answer: %d\n", day2)
	day2pt2 := day2pt2()
	fmt.Printf("Day 2 part 2 answer: %d\n", day2pt2)

	day3 := day3()
	fmt.Printf("Day 3 answer: %d\n", day3)
	day3pt2 := day3pt2()
	fmt.Printf("Day 3 part 2 answer: %d\n", day3pt2)

	day4 := day4()
	fmt.Printf("Day 4 answer: %d\n", day4)
	day4pt2 := day4pt2()
	fmt.Printf("Day 4 part 2 answer: %d\n", day4pt2)

	day5 := day5.Day5()
	fmt.Printf("Day 5 answer: %d\n", day5)
}
