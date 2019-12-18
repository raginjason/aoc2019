package intcode

import (
	"github.com/google/go-cmp/cmp"
	"strconv"
	"testing"
)

func TestAddition(t *testing.T) {
	testCases := []struct {
		testName               string
		program                Program
		wantProgram            Program
		wantInstructionPointer int
	}{
		{"position-mode add", Program{1, 0, 0, 0, 99}, Program{2, 0, 0, 0, 99}, 5},
		{"store add post terminate", Program{1, 4, 4, 5, 99, 0}, Program{1, 4, 4, 5, 99, 198}, 5},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for {
				if _, ok := <-out; !ok {
					break
				}
			}

			if diff := cmp.Diff(tc.wantProgram, c.Program); diff != "" {
				t.Errorf("Program difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	testCases := []struct {
		testName               string
		program                Program
		wantProgram            Program
		wantInstructionPointer int
	}{
		{"position-mode multiply", Program{2, 3, 0, 3, 99}, Program{2, 3, 0, 6, 99}, 5},
		{"store multiply post terminate", Program{2, 4, 4, 5, 99, 0}, Program{2, 4, 4, 5, 99, 9801}, 5},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for {
				if _, ok := <-out; !ok {
					break
				}
			}

			if diff := cmp.Diff(tc.wantProgram, c.Program); diff != "" {
				t.Errorf("Program difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestAdditionMultiplication(t *testing.T) {
	testCases := []struct {
		testName               string
		program                Program
		wantProgram            Program
		wantInstructionPointer int
	}{
		{"day 2", Program{1, 1, 1, 4, 99, 5, 6, 0, 99}, Program{30, 1, 1, 4, 2, 5, 6, 0, 99}, 9},
	}

	for _, tc := range testCases {
		t.Run(tc.program.String(), func(t *testing.T) {
			c := NewComputer(nil, tc.program, nil)
			c.Run()

			if diff := cmp.Diff(tc.wantProgram, c.Program); diff != "" {
				t.Errorf("Program difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestTerminate(t *testing.T) {
	testCases := []struct {
		testName               string
		program                Program
		wantProgram            Program
		wantInstructionPointer int
	}{
		{"only terminate", Program{99}, Program{99}, 1},
		{"terminate before end", Program{99, 1, 2, 3}, Program{99, 1, 2, 3}, 1},
		{"overwrite terminate", Program{1, 1, 1, 4, 99, 5, 6, 0, 99}, Program{30, 1, 1, 4, 2, 5, 6, 0, 99}, 9},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for {
				if _, ok := <-out; !ok {
					break
				}
			}

			if diff := cmp.Diff(tc.wantProgram, c.Program); diff != "" {
				t.Errorf("Program difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestInput(t *testing.T) {
	testCases := []struct {
		testName               string
		inputData              []int
		program                Program
		wantProgram            Program
		wantInstructionPointer int
	}{
		{"5 input to address 0", []int{5}, Program{3, 0, 99}, Program{5, 0, 99}, 3},
		{"8 input to address beyond termination", []int{8}, Program{3, 3, 99, 0}, Program{3, 3, 99, 8}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for _, v := range tc.inputData {
				in <- v
			}

			var got []int
			for {
				val, ok := <-out
				if ok == false {
					break
				} else {
					got = append(got, val)
				}
			}

			if diff := cmp.Diff(tc.wantProgram, c.Program); diff != "" {
				t.Errorf("Program difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	testCases := []struct {
		testName               string
		program                Program
		wantOutput             []int
		wantInstructionPointer int
	}{
		{"address 0", Program{4, 0, 99}, []int{4}, 3},
		{"address 2", Program{4, 2, 99}, []int{99}, 3},
		{"address beyond termination", Program{4, 3, 99, 23}, []int{23}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			var got []int
			for {
				val, ok := <-out
				if ok == false {
					break
				} else {
					got = append(got, val)
				}
			}

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("Output difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestInputOutput(t *testing.T) {
	testCases := []struct {
		testName               string
		inputData              []int
		program                Program
		wantProgram            Program
		wantOutput             []int
		wantInstructionPointer int
	}{
		// The program 3,0,4,0,99 outputs whatever it gets as input, then halts.
		{"5 in 5 out", []int{5}, Program{3, 0, 4, 0, 99}, Program{5, 0, 4, 0, 99}, []int{5}, 5},
		{"7 in 7 out", []int{7}, Program{3, 0, 4, 0, 99}, Program{7, 0, 4, 0, 99}, []int{7}, 5},
		{"1 in 1 out", []int{1}, Program{3, 0, 4, 0, 99}, Program{1, 0, 4, 0, 99}, []int{1}, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for _, v := range tc.inputData {
				in <- v
			}

			var got []int
			for {
				val, ok := <-out
				if ok == false {
					break
				} else {
					got = append(got, val)
				}
			}

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("Output difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantProgram, c.Program); diff != "" {
				t.Errorf("Program difference, -want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantInstructionPointer, c.InstructionPointer); diff != "" {
				t.Errorf("Instruction pointer difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestInstruction(t *testing.T) {
	testCases := []struct {
		input int
		want  *Instruction
	}{
		{99, &Instruction{Terminate, []ParameterMode{}}},

		{1, &Instruction{Add, []ParameterMode{PositionMode, PositionMode, PositionMode}}},
		{101, &Instruction{Add, []ParameterMode{ImmediateMode, PositionMode, PositionMode}}},
		{1001, &Instruction{Add, []ParameterMode{PositionMode, ImmediateMode, PositionMode}}},
		{1101, &Instruction{Add, []ParameterMode{ImmediateMode, ImmediateMode, PositionMode}}},
		{10001, &Instruction{Add, []ParameterMode{PositionMode, PositionMode, ImmediateMode}}},
		{10101, &Instruction{Add, []ParameterMode{ImmediateMode, PositionMode, ImmediateMode}}},
		{11001, &Instruction{Add, []ParameterMode{PositionMode, ImmediateMode, ImmediateMode}}},
		{11101, &Instruction{Add, []ParameterMode{ImmediateMode, ImmediateMode, ImmediateMode}}},

		{2, &Instruction{Multiply, []ParameterMode{PositionMode, PositionMode, PositionMode}}},
		{102, &Instruction{Multiply, []ParameterMode{ImmediateMode, PositionMode, PositionMode}}},
		{1002, &Instruction{Multiply, []ParameterMode{PositionMode, ImmediateMode, PositionMode}}},
		{1102, &Instruction{Multiply, []ParameterMode{ImmediateMode, ImmediateMode, PositionMode}}},
		{10002, &Instruction{Multiply, []ParameterMode{PositionMode, PositionMode, ImmediateMode}}},
		{10102, &Instruction{Multiply, []ParameterMode{ImmediateMode, PositionMode, ImmediateMode}}},
		{11002, &Instruction{Multiply, []ParameterMode{PositionMode, ImmediateMode, ImmediateMode}}},
		{11102, &Instruction{Multiply, []ParameterMode{ImmediateMode, ImmediateMode, ImmediateMode}}},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			ins := NewInstruction(tc.input)

			if diff := cmp.Diff(tc.want, ins); diff != "" {
				t.Errorf("Instruction difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestJump(t *testing.T) {
	bigProgram := Program{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	testCases := []struct {
		inputData  []int
		program    Program
		wantOutput []int
	}{
		// Output 0 if input was 0 else output 1, mode=position
		{[]int{-1}, Program{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{1}},
		{[]int{0}, Program{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{0}},
		{[]int{1}, Program{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{1}},
		{[]int{2}, Program{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{1}},

		// Output 0 if input was 0 else output 1, mode=immediate
		{[]int{-1}, Program{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{1}},
		{[]int{0}, Program{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{0}},
		{[]int{1}, Program{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{1}},
		{[]int{2}, Program{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{1}},

		/*
		 * bigProgram uses an input instruction to ask for a single number. The program will then output 999 if the
		 * input value is below 8, output 1000 if the input value is equal to 8, or output 1001 if the input value is
		 * greater than 8.
		 */
		{[]int{7}, bigProgram, []int{999}},
		{[]int{8}, bigProgram, []int{1000}},
		{[]int{9}, bigProgram, []int{1001}},
	}

	for _, tc := range testCases {
		t.Run(tc.program.String(), func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for _, v := range tc.inputData {
				in <- v
			}

			var got []int
			for {
				val, ok := <-out
				if ok == false {
					break
				} else {
					got = append(got, val)
				}
			}

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("Output difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	testCases := []struct {
		testName   string
		inputData  []int
		program    Program
		wantOutput []int
	}{
		{"position-mode test input 7 equal to 8", []int{7}, Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}},
		{"position-mode test input 8 equal to 8", []int{8}, Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{1}},
		{"position-mode test input 9 equal to 8", []int{9}, Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}},

		{"immediate-mode test input 7 equal to 8", []int{7}, Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{0}},
		{"immediate-mode test input 8 equal to 8", []int{8}, Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{1}},
		{"immediate-mode test input 9 equal to 8", []int{9}, Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{0}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for _, v := range tc.inputData {
				in <- v
			}

			var got []int
			for {
				val, ok := <-out
				if ok == false {
					break
				} else {
					got = append(got, val)
				}
			}

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("Output difference, -want +got:\n%s", diff)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	testCases := []struct {
		testName   string
		inputData  []int
		program    Program
		wantOutput []int
	}{
		{"position-mode test input 7 less than 8", []int{7}, Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{1}},
		{"position-mode test input 8 less than 8", []int{8}, Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}},
		{"position-mode test input 9 less than 8", []int{9}, Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}},

		{"immediate-mode test input 7 less than 8", []int{7}, Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{1}},
		{"immediate-mode test input 8 less than 8", []int{8}, Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{0}},
		{"immediate-mode test input 9 less than 8", []int{9}, Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{0}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			for _, v := range tc.inputData {
				in <- v
			}

			var got []int
			for {
				val, ok := <-out
				if ok == false {
					break
				} else {
					got = append(got, val)
				}
			}

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("Output difference, -want +got:\n%s", diff)
			}
		})
	}
}