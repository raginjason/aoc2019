package main

import (
	"fmt"
	"github.com/raginjason/aoc2019/day2"
	"github.com/raginjason/aoc2019/day5"
	"github.com/raginjason/aoc2019/day7"
	"github.com/raginjason/aoc2019/day9"
)

func main() {
	day1 := day1()
	fmt.Printf("Day 1 answer: %d\n", day1)
	day1pt2 := day1pt2()
	fmt.Printf("Day 1 part 2 answer: %d\n", day1pt2)

	day2pt1 := day2.Part1()
	fmt.Printf("Day 2 part 1 answer: %d\n", day2pt1)
	day2pt2 := day2.Part2()
	fmt.Printf("Day 2 part 2 answer: %d\n", day2pt2)

	day3 := day3()
	fmt.Printf("Day 3 answer: %d\n", day3)
	day3pt2 := day3pt2()
	fmt.Printf("Day 3 part 2 answer: %d\n", day3pt2)

	day4 := day4()
	fmt.Printf("Day 4 answer: %d\n", day4)
	day4pt2 := day4pt2()
	fmt.Printf("Day 4 part 2 answer: %d\n", day4pt2)

	day5pt1 := day5.Part1()
	fmt.Printf("Day 5 part 1 answer: %d\n", day5pt1)
	day5pt2 := day5.Part2()
	fmt.Printf("Day 5 part 2 answer: %d\n", day5pt2)

	day7pt1 := day7.Part1()
	fmt.Printf("Day 7 part 1 answer: %d\n", day7pt1)
	day7pt2 := day7.Part2()
	fmt.Printf("Day 7 part 2 answer: %d\n", day7pt2)

	day9pt1 := day9.Part1()
	fmt.Printf("Day 9 part 1 answer: %d\n", day9pt1)
	day9pt2 := day9.Part2()
	fmt.Printf("Day 9 part 2 answer: %d\n", day9pt2)
}
