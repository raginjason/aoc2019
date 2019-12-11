package day5

import (
	"github.com/raginjason/aoc2019/intcode"
	"github.com/raginjason/aoc2019/util"
	"log"
	"os"
)

func scanDay5File() []int {
	fh, err := os.Open(util.InputFilePath(5))
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer fh.Close()

	program := intcode.ParseIntCode(fh)

	return program
}

func Part1() int {
	program := scanDay5File()

	c := intcode.NewComputer([]int{1}, program)

	output := c.Run()
	return output[len(output)-1]
}

func Part2() int {
	program := scanDay5File()

	c := intcode.NewComputer([]int{5}, program)

	output := c.Run()
	return output[0]
}
