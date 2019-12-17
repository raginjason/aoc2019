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

	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(in, program, out)
	go c.Run()
	in <- 1

	var output []int
	for {
		val, ok := <-out
		if !ok {
			break
		} else {
			output = append(output, val)
		}
	}
	return output[len(output)-1]
}

func Part2() int {
	program := scanDay5File()

	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(in, program, out)
	go c.Run()
	in <- 5

	var output []int
	for {
		val, ok := <-out
		if !ok {
			break
		} else {
			output = append(output, val)
		}
	}
	return output[0]
}
