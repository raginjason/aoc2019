package day7

import (
	"github.com/google/go-cmp/cmp"
	"github.com/raginjason/aoc2019/intcode"
	"testing"
)

func TestAmplifier(t *testing.T) {
	testCases := []struct {
		signalSeq  []int
		program    intcode.Program
		wantSignal int
	}{
		{[]int{4, 3, 2, 1, 0}, intcode.Program{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
		{[]int{0, 1, 2, 3, 4}, intcode.Program{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
		{[]int{1, 0, 4, 3, 2}, intcode.Program{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, 65210},
	}

	for _, tc := range testCases {
		t.Run(tc.program.String(), func(t *testing.T) {
			got := Amplifier(tc.program, tc.signalSeq)

			if diff := cmp.Diff(tc.wantSignal, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestMaxSignal(t *testing.T) {
	testCases := []struct {
		program    intcode.Program
		wantSignal int
	}{
		{intcode.Program{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
		{intcode.Program{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
		{intcode.Program{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, 65210},
	}

	for _, tc := range testCases {
		t.Run(tc.program.String(), func(t *testing.T) {
			got := MaxSignal(tc.program)

			if diff := cmp.Diff(tc.wantSignal, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
