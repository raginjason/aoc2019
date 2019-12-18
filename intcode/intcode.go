package intcode

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

type ParameterMode int

const (
	PositionMode ParameterMode = iota
	ImmediateMode
)

func (m ParameterMode) String() string {
	if m < PositionMode || m > ImmediateMode {
		return "Unknown"
	}

	return [...]string{"Position", "Immediate"}[m]
}

type OpCode int

const (
	Add OpCode = iota + 1
	Multiply
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	Terminate OpCode = 99
)

func (op OpCode) String() string {
	if op < Add || op > Terminate {
		return "Unknown"
	}

	if op == Terminate {
		return "Terminate"
	}

	names := [...]string{
		"N/A",
		"Add",
		"Multiply",
		"Input",
		"Output",
		"JumpIfTrue",
		"JumpIfFalse",
		"LessThan",
		"Equals",
	}

	return names[op]
}

type Instruction struct {
	OpCode         OpCode
	ParameterModes []ParameterMode
}

func NewInstruction(i int) *Instruction {
	ins := new(Instruction)

	str := strconv.Itoa(i)
	letters := strings.Split(str, "")

	var opStr string

	if len(letters) > 1 {
		opStr = letters[len(letters)-2] + letters[len(letters)-1]
		letters = letters[:len(letters)-2] // Remove last 2 items
	} else if len(letters) == 1 {
		opStr = letters[len(letters)-1]
		letters = letters[:len(letters)-1] // Remove last item
	}

	// Setup the OpCode
	if val, err := strconv.Atoi(opStr); err == nil {
		ins.OpCode = OpCode(val)
	} else {
		log.Fatalf("could not convert %s%s to int OpCode: %s", letters[len(letters)-2], letters[len(letters)-1], err)
	}

	var parms int
	switch ins.OpCode {
	case Add, Multiply:
		parms = 3
	case Input, Output:
		parms = 1
	case JumpIfTrue, JumpIfFalse:
		parms = 2
	case LessThan, Equals:
		parms = 3
	default:
		parms = 0
	}

	ins.ParameterModes = make([]ParameterMode, parms)

	// Reverse in place
	for i := len(letters)/2 - 1; i >= 0; i-- {
		opp := len(letters) - 1 - i
		letters[i], letters[opp] = letters[opp], letters[i]
	}

	for i := range letters {
		if val, err := strconv.Atoi(letters[i]); err == nil {
			ins.ParameterModes[i] = ParameterMode(val)
		} else {
			log.Fatalf("could not convert value %s to int ParameterMode[%d]: %s", letters[i], i, err)
		}
	}

	return ins
}

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
	for i := 0; i < len(c.Program); {

		op := c.Program[i]

		if op == 99 { // Terminate
			if c.output != nil {
				close(c.output)
			}
			c.InstructionPointer++
			break
		}

		ins := NewInstruction(op)
		switch ins.OpCode {
		case Add:
			var input1Value int
			var input2Value int

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				input1Value = c.Program[i+1]
			case PositionMode:
				input1Value = c.Program[c.Program[i+1]]
			}

			switch ins.ParameterModes[1] {
			case ImmediateMode:
				input2Value = c.Program[i+2]
			case PositionMode:
				input2Value = c.Program[c.Program[i+2]]
			}

			outputAddress := c.Program[i+3]
			c.Program[outputAddress] = input1Value + input2Value
			i = i + 4
		case Multiply:
			var input1Value int
			var input2Value int

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				input1Value = c.Program[i+1]
			case PositionMode:
				input1Value = c.Program[c.Program[i+1]]
			}

			switch ins.ParameterModes[1] {
			case ImmediateMode:
				input2Value = c.Program[i+2]
			case PositionMode:
				input2Value = c.Program[c.Program[i+2]]
			}

			outputAddress := c.Program[i+3]
			c.Program[outputAddress] = input1Value * input2Value
			i = i + 4
		case Input:
			/*
			 * Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For
			 * example, the instruction 3,50 would take an input value and store it at address 50.
			 */
			var outputAddress int

			switch ins.ParameterModes[0] {
			case PositionMode: // zeroval
				outputAddress = c.Program[i+1]
			}

			if val, ok := <-c.input; ok {
				c.Program[outputAddress] = val
			}

			i = i + 2
		case Output:
			/*
			 * Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the
			 * value at address 50.
			 */
			var outputValue int

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				outputValue = c.Program[i+1]
			case PositionMode:
				outputValue = c.Program[c.Program[i+1]]
			}

			c.output <- outputValue
			i = i + 2
		case JumpIfTrue:
			/*
			 * If the first parameter is non-zero, it sets the instruction pointer to the value from the second
			 * parameter. Otherwise, it does nothing.
			 */
			var flag int
			var newPointer int

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				flag = c.Program[i+1]
			case PositionMode:
				flag = c.Program[c.Program[i+1]]
			}

			switch ins.ParameterModes[1] {
			case ImmediateMode:
				newPointer = c.Program[i+2]
			case PositionMode:
				newPointer = c.Program[c.Program[i+2]]
			}

			if flag != 0 {
				i = newPointer
			} else {
				i = i + 3
			}
		case JumpIfFalse:
			/*
			 * If the first parameter is zero, it sets the instruction pointer to the value from the second parameter.
			 * Otherwise, it does nothing.
			 */
			var flag int
			var newPointer int

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				flag = c.Program[i+1]
			case PositionMode:
				flag = c.Program[c.Program[i+1]]
			}

			switch ins.ParameterModes[1] {
			case ImmediateMode:
				newPointer = c.Program[i+2]
			case PositionMode:
				newPointer = c.Program[c.Program[i+2]]
			}

			if flag == 0 {
				i = newPointer
			} else {
				i = i + 3
			}
		case LessThan:
			/*
			 * If the first parameter is less than the second parameter, it stores 1 in the position given by the third
			 * parameter. Otherwise, it stores 0.
			 */

			var input1Value int
			var input2Value int
			outputAddress := c.Program[i+3]

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				input1Value = c.Program[i+1]
			case PositionMode:
				input1Value = c.Program[c.Program[i+1]]
			}

			switch ins.ParameterModes[1] {
			case ImmediateMode:
				input2Value = c.Program[i+2]
			case PositionMode:
				input2Value = c.Program[c.Program[i+2]]
			}

			if input1Value < input2Value {
				c.Program[outputAddress] = 1
			} else {
				c.Program[outputAddress] = 0
			}
			i = i + 4
		case Equals:
			/*
			 * Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position
			 * given by the third parameter. Otherwise, it stores 0.
			 */

			var input1Value int
			var input2Value int
			outputAddress := c.Program[i+3]

			switch ins.ParameterModes[0] {
			case ImmediateMode:
				input1Value = c.Program[i+1]
			case PositionMode:
				input1Value = c.Program[c.Program[i+1]]
			}

			switch ins.ParameterModes[1] {
			case ImmediateMode:
				input2Value = c.Program[i+2]
			case PositionMode:
				input2Value = c.Program[c.Program[i+2]]
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
		c.InstructionPointer = i
	}
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
