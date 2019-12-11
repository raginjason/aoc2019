package day5

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestInputOutput(t *testing.T) {
	testCases := []struct {
		inputData      []int
		initialProgram program
		finalProgram   program
		wantOutput     []int
	}{
		// The program 3,0,4,0,99 outputs whatever it gets as input, then halts.
		{[]int{5}, program{3, 0, 4, 0, 99}, program{5, 0, 4, 0, 99}, []int{5}},
		{[]int{7}, program{3, 0, 4, 0, 99}, program{7, 0, 4, 0, 99}, []int{7}},
		{[]int{1}, program{3, 0, 4, 0, 99}, program{1, 0, 4, 0, 99}, []int{1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.initialProgram), func(t *testing.T) {
			c := NewComputer(tc.inputData, tc.initialProgram)
			got := c.Run()

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.finalProgram, c.program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestPositionMode(t *testing.T) {
	testCases := []struct {
		inputData      []int
		initialProgram program
		finalProgram   program
		wantOutput     []int
	}{
		{nil, program{1, 0, 0, 0, 99}, program{2, 0, 0, 0, 99}, nil}, // Add
		{nil, program{2, 3, 0, 3, 99}, program{2, 3, 0, 6, 99}, nil}, // Multiply
		{nil, program{2, 4, 4, 5, 99, 0}, program{2, 4, 4, 5, 99, 9801}, nil},
		{nil, program{1, 1, 1, 4, 99, 5, 6, 0, 99}, program{30, 1, 1, 4, 2, 5, 6, 0, 99}, nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.initialProgram), func(t *testing.T) {
			c := NewComputer(tc.inputData, tc.initialProgram)
			got := c.Run()

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.finalProgram, c.program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestImmediateMode(t *testing.T) {
	testCases := []struct {
		inputData      []int
		initialProgram program
		finalProgram   program
		wantOutput     []int
	}{
		{nil, program{1002, 4, 3, 4, 33}, program{1002, 4, 3, 4, 99}, nil},      // Multiply
		{nil, program{1101, 100, -1, 4, 0}, program{1101, 100, -1, 4, 99}, nil}, // Add
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.initialProgram), func(t *testing.T) {
			c := NewComputer(tc.inputData, tc.initialProgram)
			got := c.Run()

			if diff := cmp.Diff(tc.wantOutput, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.finalProgram, c.program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
