package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestFuelCalc(t *testing.T) {
	testCases := []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d->%d", tc.mass, tc.want), func(t *testing.T) {
			got := ComputeFuel(tc.mass)
			if tc.want != got {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestFuelTotalCalc(t *testing.T) {
	testCases := []struct {
		list []int
		want int
	}{
		{[]int{12, 12, 1969, 100756}, 34241},
	}

	for _, tc := range testCases {
		var strs []string
		for _, num := range tc.list {
			s := strconv.Itoa(num)
			strs = append(strs, s)
		}
		name := strings.Join(strs, "+")

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s->%d", name, tc.want), func(t *testing.T) {
				got := ComputeFuelTotal(tc.list)
				if tc.want != got {
					t.Errorf("got %d; want %d", got, tc.want)
				}
			})
		}
	}
}
