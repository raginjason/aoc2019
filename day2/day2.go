package day2

import (
	"github.com/raginjason/aoc2019/intcode"
	"github.com/raginjason/aoc2019/util"
	"log"
	"os"
)

func scanDay2File() []int {
	fh, err := os.Open(util.InputFilePath(2))
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer fh.Close()

	program := intcode.ParseIntCode(fh)

	return program
}

func Part1() int {
	program := scanDay2File()

	c := intcode.NewComputer(nil, program)
	// replace position 1 with the value 12 and replace position 2 with the value 2
	c.Program[1] = 12
	c.Program[2] = 2

	c.Run()
	return c.Program[0]
}

func Part2() int {
	opCodes := scanDay2File()

	for noun := 1; noun <= 99; noun++ {
		for verb := 1; verb <= 99; verb++ {
			newOpCodes := make([]int, len(opCodes))
			copy(newOpCodes, opCodes)
			newOpCodes[1] = noun
			newOpCodes[2] = verb

			c := intcode.NewComputer(nil, newOpCodes)
			c.Run()

			if c.Program[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}
