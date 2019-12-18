package intcode

import (
	"github.com/google/go-cmp/cmp"
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
