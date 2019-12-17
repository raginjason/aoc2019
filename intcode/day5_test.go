package intcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestInputOutput(t *testing.T) {
	testCases := []struct {
		inputData      []int
		initialProgram Program
		finalProgram   Program
		wantOutput     []int
	}{
		// The program 3,0,4,0,99 outputs whatever it gets as input, then halts.
		{[]int{5}, Program{3, 0, 4, 0, 99}, Program{5, 0, 4, 0, 99}, []int{5}},
		{[]int{7}, Program{3, 0, 4, 0, 99}, Program{7, 0, 4, 0, 99}, []int{7}},
		{[]int{1}, Program{3, 0, 4, 0, 99}, Program{1, 0, 4, 0, 99}, []int{1}},
	}

	for _, tc := range testCases {
		t.Run(tc.initialProgram.String(), func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.initialProgram, out)
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
				t.Errorf("-want +got:\n%s", diff)
			}

			if diff := cmp.Diff(tc.finalProgram, c.Program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestPositionMode(t *testing.T) {
	testCases := []struct {
		inputData      []int
		initialProgram Program
		finalProgram   Program
		wantOutput     []int
	}{
		{nil, Program{1, 0, 0, 0, 99}, Program{2, 0, 0, 0, 99}, nil}, // Add
		{nil, Program{2, 3, 0, 3, 99}, Program{2, 3, 0, 6, 99}, nil}, // Multiply
		{nil, Program{2, 4, 4, 5, 99, 0}, Program{2, 4, 4, 5, 99, 9801}, nil},
		{nil, Program{1, 1, 1, 4, 99, 5, 6, 0, 99}, Program{30, 1, 1, 4, 2, 5, 6, 0, 99}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.initialProgram.String(), func(t *testing.T) {
			c := NewComputer(nil, tc.initialProgram, nil)
			c.Run()

			if diff := cmp.Diff(tc.finalProgram, c.Program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestImmediateMode(t *testing.T) {
	testCases := []struct {
		inputData      []int
		initialProgram Program
		finalProgram   Program
		wantOutput     []int
	}{
		{nil, Program{1002, 4, 3, 4, 33}, Program{1002, 4, 3, 4, 99}, nil},      // Multiply
		{nil, Program{1101, 100, -1, 4, 0}, Program{1101, 100, -1, 4, 99}, nil}, // Add
	}

	for _, tc := range testCases {
		t.Run(tc.initialProgram.String(), func(t *testing.T) {
			c := NewComputer(nil, tc.initialProgram, nil)
			c.Run()

			if diff := cmp.Diff(tc.finalProgram, c.Program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		inputData  []int
		program    Program
		wantOutput []int
	}{
		// 3,9,8,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
		{[]int{7}, Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}}, // 7 equal to 8, mode=position
		{[]int{8}, Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{1}}, // 8 equal to 8, mode=position
		{[]int{9}, Program{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}}, // 9 equal to 8, mode=position

		// 3,3,1108,-1,8,3,4,3,99 - Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
		{[]int{7}, Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{0}}, // 7 equal to 8, mode=immediate
		{[]int{8}, Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{1}}, // 8 equal to 8, mode=immediate
		{[]int{9}, Program{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{0}}, // 9 equal to 8, mode=immediate
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
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	testCases := []struct {
		inputData  []int
		program    Program
		wantOutput []int
	}{
		// 3,9,7,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
		{[]int{7}, Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{1}}, // 7 less than 8, mode=position
		{[]int{8}, Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}}, // 8 less than 8, mode=position
		{[]int{9}, Program{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{0}}, // 9 less than 8, mode=position

		// 3,3,1107,-1,8,3,4,3,99 - Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
		{[]int{7}, Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{1}}, // 7 less than 8, mode=immediate
		{[]int{8}, Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{0}}, // 8 less than 8, mode=immediate
		{[]int{9}, Program{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{0}}, // 9 less than 8, mode=immediate
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
				t.Errorf("-want +got:\n%s", diff)
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
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
