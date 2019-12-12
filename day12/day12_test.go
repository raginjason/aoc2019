package day12

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

/*
<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
*/

func TestSimulation(t *testing.T) {
	testCases := []struct {
		name  string
		steps int
		input []Moon
		want  []Moon
	}{
		{"0/10", 0,
			[]Moon{
				{Position{-1, 0, 2}, Velocity{0, 0, 0},},
				{Position{2, -10, -7}, Velocity{0, 0, 0},},
				{Position{4, -8, 8}, Velocity{0, 0, 0},},
				{Position{3, 5, -1}, Velocity{0, 0, 0},},
			},
			[]Moon{
				{Position{-1, 0, 2}, Velocity{0, 0, 0},},
				{Position{2, -10, -7}, Velocity{0, 0, 0},},
				{Position{4, -8, 8}, Velocity{0, 0, 0},},
				{Position{3, 5, -1}, Velocity{0, 0, 0},},
			},
		},
		{"1/10", 1,
			[]Moon{
				{Position{-1, 0, 2}, Velocity{0, 0, 0},},
				{Position{2, -10, -7}, Velocity{0, 0, 0},},
				{Position{4, -8, 8}, Velocity{0, 0, 0},},
				{Position{3, 5, -1}, Velocity{0, 0, 0},},
			},
			[]Moon{
				{Position{2, -1, 1}, Velocity{3, -1, -1},},
				{Position{3, -7, -4}, Velocity{1, 3, 3},},
				{Position{1, -7, 5}, Velocity{-3, 1, -3},},
				{Position{2, 2, 0}, Velocity{-1, -3, 1},},
			},
		},
		{"10/10", 10,
			[]Moon{
				{Position{-1, 0, 2}, Velocity{0, 0, 0},},
				{Position{2, -10, -7}, Velocity{0, 0, 0},},
				{Position{4, -8, 8}, Velocity{0, 0, 0},},
				{Position{3, 5, -1}, Velocity{0, 0, 0},},
			},
			[]Moon{
				{Position{2, 1, -3}, Velocity{-3, -2, 1},},
				{Position{1, -8, 0}, Velocity{-1, 1, 3},},
				{Position{3, -6, 1}, Velocity{3, 2, -3},},
				{Position{2, 0, 4}, Velocity{1, -1, -1},},
			},
		},
		{"100/100", 100,
			/*
				After 0 steps:
				pos=<x= -8, y=-10, z=  0>, vel=<x=  0, y=  0, z=  0>
				pos=<x=  5, y=  5, z= 10>, vel=<x=  0, y=  0, z=  0>
				pos=<x=  2, y= -7, z=  3>, vel=<x=  0, y=  0, z=  0>
				pos=<x=  9, y= -8, z= -3>, vel=<x=  0, y=  0, z=  0>
			*/
			[]Moon{
				{Position{-8, -10, 0}, Velocity{0, 0, 0},},
				{Position{5, 5, 10}, Velocity{0, 0, 0},},
				{Position{2, -7, 3}, Velocity{0, 0, 0},},
				{Position{9, -8, -3}, Velocity{0, 0, 0},},
			},
			/*
				After 100 steps:
				pos=<x=  8, y=-12, z= -9>, vel=<x= -7, y=  3, z=  0>
				pos=<x= 13, y= 16, z= -3>, vel=<x=  3, y=-11, z= -5>
				pos=<x=-29, y=-11, z= -1>, vel=<x= -3, y=  7, z=  4>
				pos=<x= 16, y=-13, z= 23>, vel=<x=  7, y=  1, z=  1>
			*/
			[]Moon{
				{Position{8, -12, -9}, Velocity{-7, 3, 0},},
				{Position{13, 16, -3}, Velocity{3, -11, -5},},
				{Position{-29, -11, -1}, Velocity{-3, 7, 4},},
				{Position{16, -13, 23}, Velocity{7, 1, 1},},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Simulate(tc.input, tc.steps)
			if diff := cmp.Diff(tc.want, tc.input); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}

}

func TestEnergy(t *testing.T) {
	testCases := []struct {
		input         []Moon
		wantPotential []int
		wantKinetic   []int
		wantTotal     []int
	}{
		{[]Moon{
			{Position{2, 1, -3}, Velocity{-3, -2, 1},},
			{Position{1, -8, 0}, Velocity{-1, 1, 3},},
			{Position{3, -6, 1}, Velocity{3, 2, -3},},
			{Position{2, 0, 4}, Velocity{1, -1, -1},},
		},
			[]int{6, 9, 10, 6},
			[]int{6, 5, 8, 3},
			[]int{36, 45, 80, 18},
		},
		{[]Moon{
			{Position{8, -12, -9}, Velocity{-7, 3, 0},},
			{Position{13, 16, -3}, Velocity{3, 11, -5},},
			{Position{-29, 11, -1}, Velocity{-3, 7, 4},},
			{Position{16, 13, 23}, Velocity{7, 1, 1},},
		},
			[]int{29, 32, 41, 52},
			[]int{10, 19, 14, 9},
			[]int{290, 608, 574, 468},
		},
	}

	for _, tc := range testCases {
		for i := range tc.input {
			t.Run("potential", func(t *testing.T) {
				if diff := cmp.Diff(tc.wantPotential[i], tc.input[i].PotentialEnergy()); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			})
			t.Run("kinetic", func(t *testing.T) {
				if diff := cmp.Diff(tc.wantKinetic[i], tc.input[i].KineticEnergy()); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			})
			t.Run("total", func(t *testing.T) {
				if diff := cmp.Diff(tc.wantTotal[i], tc.input[i].TotalEnergy()); diff != "" {
					t.Errorf("-want +got:\n%s", diff)
				}
			})
		}
	}
}

/*
For example, if A has an x position of 3, and B has a x position of 5, then A's x velocity changes
by +1 (because 5 > 3) and B's x velocity changes by -1 (because 3 < 5). However, if the positions on a given
axis are the same, the velocity on that axis does not change for that pair of moons.
*/
func TestCalculateVelocityDelta(t *testing.T) {
	testCases := []struct {
		inputA int
		inputB int
		wantA  int
		wantB  int
	}{
		{0, 0, 0, 0},
		{3, 5, 1, -1},
		{3, 3, 0, 0},
		{8, 3, -1, 1},
		{-3, -5, -1, 1},
		{-3, -3, 0, 0},
		{-8, -3, 1, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d,%d", tc.inputA, tc.inputB), func(t *testing.T) {
			gotA, gotB := CalculateVelocityDelta(tc.inputA, tc.inputB)
			if diff := cmp.Diff(tc.wantA, gotA); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.wantB, gotB); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

//position of x=1, y=2, z=3 and a velocity of x=-2, y=0,z=3, then its new position would be x=-1, y=2, z=6
func TestMoon_ApplyVelocity(t *testing.T) {
	testCases := []struct {
		input        Moon
		wantPosition Position
	}{
		{Moon{Position{1, 2, 3}, Velocity{-2, 0, 3}}, Position{-1, 2, 6}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d,%d,%d", tc.input.Pos.X, tc.input.Pos.Y, tc.input.Pos.Z), func(t *testing.T) {
			tc.input.ApplyVelocity()
			if diff := cmp.Diff(tc.wantPosition, tc.input.Pos); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
