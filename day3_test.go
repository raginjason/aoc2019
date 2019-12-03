package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCrossedWires(t *testing.T) {
	testCases := []struct {
		inputA []string
		inputB []string
		want   int
	}{
		{
			[]string{"R8", "U5", "L5", "D3"},
			[]string{"U7", "R6", "D4", "L4"},
			6,
		},
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			159,
		},
		{
			[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			135,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%d", tc.want), func(t *testing.T) {
			got := CrossedWires(tc.inputA, tc.inputB)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

/*
I didn't feel like creating test structs for the other 2 samples provided, but this one should be good enough

UPDATE: It was not good enough, because it didn't cover going a direction more than once; need to add distance instead
of simply overwriting it. Still did not implement all of the example cases, but implemented one that revealed the bug
and fixed the logic
*/
func TestComputeSegments(t *testing.T) {
	testCases := []struct {
		input []string
		want  []Segment
	}{
		{
			[]string{"R8", "U5", "L5", "D3"},
			[]Segment{
				{true, false, Point{0, 0}, Point{8, 0}},
				{false, true, Point{8, 0}, Point{8, 5}},
				{true, false, Point{8, 5}, Point{3, 5}},
				{false, true, Point{3, 5}, Point{3, 2}},
			},
		},
		{
			[]string{"U7", "R6", "D4", "L4"},
			[]Segment{
				{false, true, Point{0, 0}, Point{0, 7}},
				{true, false, Point{0, 7}, Point{6, 7}},
				{false, true, Point{6, 7}, Point{6, 3}},
				{true, false, Point{6, 3}, Point{2, 3}},
			},
		},
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]Segment{
				{true, false, Point{0, 0}, Point{75, 0}},
				{false, true, Point{75, 0}, Point{75, -30}},
				{true, false, Point{75, -30}, Point{158, -30}},
				{false, true, Point{158, -30}, Point{158, 53}},
				{true, false, Point{158, 53}, Point{146, 53}},
				{false, true, Point{146, 53}, Point{146, 4}},
				{true, false, Point{146, 4}, Point{217, 4}},
				{false, true, Point{217, 4}, Point{217, 11}},
				{true, false, Point{217, 11}, Point{145, 11}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%s", tc.want), func(t *testing.T) {
			got := ComputeSegments(tc.input)
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(Segment{}, Point{})); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

// Made up simple segments
func TestCalculateIntersection(t *testing.T) {
	testCases := []struct {
		inputA Segment
		inputB Segment
		want   Point
	}{
		{
			Segment{false, true, Point{2, 5}, Point{2, 2}},
			Segment{true, false, Point{1, 4}, Point{8, 4}},
			Point{2, 4},
		},
		{ // Inverse order of previous
			Segment{true, false, Point{1, 4}, Point{8, 4}},
			Segment{false, true, Point{2, 5}, Point{2, 2}},
			Point{2, 4},
		},
		{ // Negative of previous
			Segment{true, false, Point{-1, -4}, Point{-8, -4}},
			Segment{false, true, Point{-2, -5}, Point{-2, -2}},
			Point{-2, -4},
		},
		{ // Crosses positive/negative boundaries
			Segment{false, true, Point{3, 4}, Point{3, -8}},
			Segment{true, false, Point{-9, -7}, Point{6, -7}},
			Point{3, -7},
		},
		{ // Inverse order of previous
			Segment{true, false, Point{-9, -7}, Point{6, -7}},
			Segment{false, true, Point{3, 4}, Point{3, -8}},
			Point{3, -7},
		},
		{ // Negative of previous
			Segment{true, false, Point{9, 7}, Point{-6, 7}},
			Segment{false, true, Point{-3, -4}, Point{-3, 8}},
			Point{-3, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%s", tc.want), func(t *testing.T) {
			got := CalculateIntersection(tc.inputA, tc.inputB)
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(Segment{}, Point{})); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
