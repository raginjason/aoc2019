package day5

import (
	"encoding/csv"
	"github.com/raginjason/aoc2019/util"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func (p program) String() string {
	var strs []string
	for _, num := range p {
		s := strconv.Itoa(num)
		strs = append(strs, s)
	}
	return strings.Join(strs, "-")
}

type program []int

type Computer struct {
	inputData []int
	program   program
	outputs   []int
}

func NewComputer(input []int, program []int) *Computer {
	c := new(Computer)
	c.inputData = input
	c.program = program
	return c
}

func (c *Computer) Run() []int {
	skip := 0
	inputCounter := 0
	for i, op := range c.program {
		if skip > 1 {
			skip = skip - 1
			continue
		}

		if op == 99 { // Terminate
			break
		}

		opCodeParts := splitInt(op)

		switch opCodeParts[len(opCodeParts)-1] {
		case 1: // Add
			var input1Value int
			var input2Value int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				input1Value = c.program[i+1] // Immediate mode
			} else {
				input1Value = c.program[c.program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				input2Value = c.program[i+2] // Immediate mode
			} else {
				input2Value = c.program[c.program[i+2]] // Position mode
			}

			outputAddress := c.program[i+3]
			c.program[outputAddress] = input1Value + input2Value
			skip = 4
		case 2: // Multiply
			var input1Value int
			var input2Value int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				input1Value = c.program[i+1] // Immediate mode
			} else {
				input1Value = c.program[c.program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				input2Value = c.program[i+2] // Immediate mode
			} else {
				input2Value = c.program[c.program[i+2]] // Position mode
			}

			outputAddress := c.program[i+3]
			c.program[outputAddress] = input1Value * input2Value
			skip = 4
		case 3:
			/*
			 * Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For
			 * example, the instruction 3,50 would take an input value and store it at address 50.
			 */
			outputAddress := c.program[i+1]
			c.program[outputAddress] = c.inputData[inputCounter]
			inputCounter = inputCounter + 1
			skip = 2
		case 4:
			/*
			 * Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the
			 * value at address 50.
			 */
			outputAddress := c.program[i+1]
			c.outputs = append(c.outputs, c.program[outputAddress])
			skip = 2
		default:
			continue
		}
	}
	return c.outputs
}

func splitInt(i int) []int {
	str := strconv.Itoa(i)
	letters := strings.Split(str, "")
	ints := []int{}
	for _, i := range letters {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalf("could not convert %s to int", i)
		}
		ints = append(ints, j)
	}
	return ints
}

func scanDay5File() []int {
	fh, err := os.Open(util.InputFilePath(5))
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer fh.Close()

	program := parseDay5Data(fh)

	return program
}

func Day5() int {
	program := scanDay5File()

	c := NewComputer([]int{1}, program)

	output := c.Run()
	return output[len(output)-1]
}

func parseDay5Data(reader io.Reader) program {
	var program = make(program, 0)

	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read from file: %s", err)
		}

		for i := range record {
			num, err := strconv.Atoi(record[i])
			if err != nil {
				log.Fatalf("could not convert %s, %s\n", record[i], err)
			}
			program = append(program, num)
		}
	}

	return program
}
