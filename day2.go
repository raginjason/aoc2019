package main

import (
	"encoding/csv"
	"github.com/raginjason/aoc2019/util"
	"io"
	"log"
	"strconv"
	"strings"
)

func GravityComputer(opCodes []int) []int {
	skip := 0
	for i, op := range opCodes {
		if skip > 1 {
			skip = skip - 1
			continue
		}

		switch op {
		case 1: // Add
			opCodes[opCodes[i+3]] = opCodes[opCodes[i+1]] + opCodes[opCodes[i+2]]
			skip = 4
		case 2: // Multiply
			opCodes[opCodes[i+3]] = opCodes[opCodes[i+1]] * opCodes[opCodes[i+2]]
			skip = 4
		case 99: // Terminate
			return opCodes
		default:
			continue
		}
	}
	return opCodes
}

func scanDay2File() []int {
	lines, err := util.FileInput(2)
	if err != nil {
		log.Fatalf("failed to get data, %s", err)
	}

	var opCodes []int

	for _, line := range lines {
		r := csv.NewReader(strings.NewReader(line))
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			for _, opCode := range record {
				num, err := strconv.Atoi(opCode)
				if err != nil {
					log.Fatalf("could not convert %s, %s\n", opCode, err)
				}
				opCodes = append(opCodes, num)
			}
		}
	}
	return opCodes
}

func day2() int {
	opCodes := scanDay2File()

	// replace position 1 with the value 12 and replace position 2 with the value 2
	opCodes[1] = 12
	opCodes[2] = 2

	res := GravityComputer(opCodes)

	return res[0]
}

func day2pt2() int {
	opCodes := scanDay2File()

	for noun := 1; noun <= 99; noun++ {
		for verb := 1; verb <= 99; verb++ {
			newOpCodes := make([]int, len(opCodes))
			copy(newOpCodes, opCodes)
			newOpCodes[1] = noun
			newOpCodes[2] = verb

			res := GravityComputer(newOpCodes)

			if res[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}
