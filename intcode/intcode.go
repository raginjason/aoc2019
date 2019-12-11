package intcode

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

func (p Program) String() string {
	var strs []string
	for _, num := range p {
		s := strconv.Itoa(num)
		strs = append(strs, s)
	}
	return strings.Join(strs, "-")
}

type Program []int

type Computer struct {
	input              chan int
	Program            Program
	InstructionPointer int
	output             chan int
	Name               string
}

func NewComputer(input chan int, program []int, output chan int) *Computer {
	c := new(Computer)
	c.input = input
	c.Program = program
	c.output = output
	c.Name = "default"
	return c
}

func (c *Computer) Run() {
	inputCounter := 0
	for i := 0; i < len(c.Program); {

		op := c.Program[i]

		if op == 99 { // Terminate
			if c.output != nil {
				close(c.output)
			}
			break
		}

		opCodeParts := splitInt(op)

		switch opCodeParts[len(opCodeParts)-1] {
		case 1: // Add
			var input1Value int
			var input2Value int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				input1Value = c.Program[i+1] // Immediate mode
			} else {
				input1Value = c.Program[c.Program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				input2Value = c.Program[i+2] // Immediate mode
			} else {
				input2Value = c.Program[c.Program[i+2]] // Position mode
			}

			outputAddress := c.Program[i+3]
			c.Program[outputAddress] = input1Value + input2Value
			i = i + 4
		case 2: // Multiply
			var input1Value int
			var input2Value int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				input1Value = c.Program[i+1] // Immediate mode
			} else {
				input1Value = c.Program[c.Program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				input2Value = c.Program[i+2] // Immediate mode
			} else {
				input2Value = c.Program[c.Program[i+2]] // Position mode
			}

			outputAddress := c.Program[i+3]
			c.Program[outputAddress] = input1Value * input2Value
			i = i + 4
		case 3:
			/*
			 * Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For
			 * example, the instruction 3,50 would take an input value and store it at address 50.
			 */
			outputAddress := c.Program[i+1]

			if val, ok := <-c.input; ok {
				c.Program[outputAddress] = val
			}

			inputCounter = inputCounter + 1
			i = i + 2
		case 4:
			/*
			 * Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the
			 * value at address 50.
			 */
			var outputValue int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				outputValue = c.Program[i+1] // Immediate mode
			} else {
				outputValue = c.Program[c.Program[i+1]] // Position mode
			}

			c.output <- outputValue
			i = i + 2
		case 5: // Jump-if-true
			/*
			 * If the first parameter is non-zero, it sets the instruction pointer to the value from the second
			 * parameter. Otherwise, it does nothing.
			 */
			var flag int
			var newPointer int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				flag = c.Program[i+1] // Immediate mode
			} else {
				flag = c.Program[c.Program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				newPointer = c.Program[i+2] // Immediate mode
			} else {
				newPointer = c.Program[c.Program[i+2]] // Position mode
			}

			if flag != 0 {
				i = newPointer
			} else {
				i = i + 3
			}
		case 6: // Jump-if-false
			/*
			 * If the first parameter is zero, it sets the instruction pointer to the value from the second parameter.
			 * Otherwise, it does nothing.
			 */
			var flag int
			var newPointer int

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				flag = c.Program[i+1] // Immediate mode
			} else {
				flag = c.Program[c.Program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				newPointer = c.Program[i+2] // Immediate mode
			} else {
				newPointer = c.Program[c.Program[i+2]] // Position mode
			}

			if flag == 0 {
				i = newPointer
			} else {
				i = i + 3
			}
		case 7: // Less than
			/*
			 * If the first parameter is less than the second parameter, it stores 1 in the position given by the third
			 * parameter. Otherwise, it stores 0.
			 */

			var input1Value int
			var input2Value int
			outputAddress := c.Program[i+3]

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				input1Value = c.Program[i+1] // Immediate mode
			} else {
				input1Value = c.Program[c.Program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				input2Value = c.Program[i+2] // Immediate mode
			} else {
				input2Value = c.Program[c.Program[i+2]] // Position mode
			}

			if input1Value < input2Value {
				c.Program[outputAddress] = 1
			} else {
				c.Program[outputAddress] = 0
			}
			i = i + 4
		case 8:
			/*
			 * Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position
			 * given by the third parameter. Otherwise, it stores 0.
			 */

			var input1Value int
			var input2Value int
			outputAddress := c.Program[i+3]

			if len(opCodeParts)-3 >= 0 && opCodeParts[len(opCodeParts)-3] == 1 {
				input1Value = c.Program[i+1] // Immediate mode
			} else {
				input1Value = c.Program[c.Program[i+1]] // Position mode
			}

			if len(opCodeParts)-4 >= 0 && opCodeParts[len(opCodeParts)-4] == 1 {
				input2Value = c.Program[i+2] // Immediate mode
			} else {
				input2Value = c.Program[c.Program[i+2]] // Position mode
			}

			if input1Value == input2Value {
				c.Program[outputAddress] = 1
			} else {
				c.Program[outputAddress] = 0
			}
			i = i + 4
		default:
			i = i + 1
		}
	}
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

func ParseIntCode(reader io.Reader) Program {
	var program = make(Program, 0)

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
