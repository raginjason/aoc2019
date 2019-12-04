package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestClosestIntersection(t *testing.T) {
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
			got := ClosestIntersection(tc.inputA, tc.inputB)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestShortestIntersection(t *testing.T) {
	testCases := []struct {
		inputA []string
		inputB []string
		want   int
	}{
		{
			[]string{"R8", "U5", "L5", "D3"},
			[]string{"U7", "R6", "D4", "L4"},
			30,
		},
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			610,
		},
		{
			[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			410,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("->%d", tc.want), func(t *testing.T) {
			got := ShortestIntersections(tc.inputA, tc.inputB)
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
				{true, false, Point{0, 0, 0}, Point{8, 0, 8}},
				{false, true, Point{8, 0, 8}, Point{8, 5, 13}},
				{true, false, Point{8, 5, 13}, Point{3, 5, 18}},
				{false, true, Point{3, 5, 18}, Point{3, 2, 21}},
			},
		},
		{
			[]string{"U7", "R6", "D4", "L4"},
			[]Segment{
				{false, true, Point{0, 0, 0}, Point{0, 7, 7}},
				{true, false, Point{0, 7, 7}, Point{6, 7, 13}},
				{false, true, Point{6, 7, 13}, Point{6, 3, 17}},
				{true, false, Point{6, 3, 17}, Point{2, 3, 21}},
			},
		},
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]Segment{
				{true, false, Point{0, 0, 0}, Point{75, 0, 75}},
				{false, true, Point{75, 0, 75}, Point{75, -30, 105}},
				{true, false, Point{75, -30, 105}, Point{158, -30, 188}},
				{false, true, Point{158, -30, 188}, Point{158, 53, 271}},
				{true, false, Point{158, 53, 271}, Point{146, 53, 283}},
				{false, true, Point{146, 53, 283}, Point{146, 4, 332}},
				{true, false, Point{146, 4, 332}, Point{217, 4, 403}},
				{false, true, Point{217, 4, 403}, Point{217, 11, 410}},
				{true, false, Point{217, 11, 410}, Point{145, 11, 482}},
			},
		},
	}

	for _, tc := range testCases {
		testName := strings.Join(tc.input, "-")

		t.Run(testName, func(t *testing.T) {
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
			Segment{false, true, Point{2, 5, 0}, Point{2, 2, 0}},
			Segment{true, false, Point{1, 4, 0}, Point{8, 4, 0}},
			Point{2, 4, 2},
		},
		{ // Inverse order of previous
			Segment{true, false, Point{1, 4, 0}, Point{8, 4, 0}},
			Segment{false, true, Point{2, 5, 0}, Point{2, 2, 0}},
			Point{2, 4, 2},
		},
		{ // Negative of previous
			Segment{true, false, Point{-1, -4, 0}, Point{-8, -4, 0}},
			Segment{false, true, Point{-2, -5, 0}, Point{-2, -2, 0}},
			Point{-2, -4, 2},
		},
		{ // Crosses positive/negative boundaries
			Segment{false, true, Point{3, 4, 0}, Point{3, -8, 0}},
			Segment{true, false, Point{-9, -7, 0}, Point{6, -7, 0}},
			Point{3, -7, 23},
		},
		{ // Inverse order of previous
			Segment{true, false, Point{-9, -7, 0}, Point{6, -7, 0}},
			Segment{false, true, Point{3, 4, 0}, Point{3, -8, 0}},
			Point{3, -7, 23},
		},
		{ // Negative of previous
			Segment{true, false, Point{9, 7, 0}, Point{-6, 7, 0}},
			Segment{false, true, Point{-3, -4, 0}, Point{-3, 8, 0}},
			Point{-3, 7, 23},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s+%s", tc.inputA, tc.inputB), func(t *testing.T) {
			got := CalculateIntersection(tc.inputA, tc.inputB)
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(Segment{}, Point{})); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
