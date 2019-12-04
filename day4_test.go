package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCheckAdjacentDigitPassword(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{122345, true},
		{111111, true},
		{123789, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%d", tc.input), func(t *testing.T) {
			got := CheckAdjacentDigitPassword(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestNonDecreasingDigitPassword(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{122345, true},
		{111123, true},
		{135679, true},
		{111111, true},
		{223450, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%d", tc.input), func(t *testing.T) {
			got := NonDecreasingDigitPassword(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestSixDigitPassword(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{99999, false},
		{100000, true},
		{100001, true},
		{999999, true},
		{1000000, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%d", tc.input), func(t *testing.T) {
			got := SixDigitPassword(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestLargerGroupAdjacentDigitPassword(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},

		{122345, true},
		{111111, false},
		{123789, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%d", tc.input), func(t *testing.T) {
			got := LargerGroupAdjacentDigitPassword(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
