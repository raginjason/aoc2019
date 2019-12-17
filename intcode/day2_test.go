package intcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDay2Instructions(t *testing.T) {
	testCases := []struct {
		testName               string
		program                Program
		wantProgram            Program
		wantInstructionPointer int
	}{
		{"position-mode add", Program{1, 0, 0, 0, 99}, Program{2, 0, 0, 0, 99}, 5},
		{"position-mode multiply", Program{2, 3, 0, 3, 99}, Program{2, 3, 0, 6, 99}, 5},
		{"store multiply post terminate", Program{2, 4, 4, 5, 99, 0}, Program{2, 4, 4, 5, 99, 9801}, 5},
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
