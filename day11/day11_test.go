package day11

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRobotTurn(t *testing.T) {
	testCases := []struct {
		testName           string
		initialOrientation Direction
		turn               Turn
		wantOrientation    Direction
	}{
		{"up -> right", Up, TurnRight, Right},
		{"right -> down", Right, TurnRight, Down},
		{"down -> left", Down, TurnRight, Left},
		{"left -> up", Left, TurnRight, Up},

		{"up -> left", Up, TurnLeft, Left},
		{"right -> up", Right, TurnLeft, Up},
		{"down -> right", Down, TurnLeft, Right},
		{"left -> down", Left, TurnLeft, Down},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			r := Robot{tc.initialOrientation, Coordinate{0, 0}}

			r.Turn(tc.turn)

			if diff := cmp.Diff(tc.wantOrientation, r.Orientation); diff != "" {
				t.Errorf("Orientation difference, -want +got:\n%s", diff)
			}

		})
	}
}

func TestRobotMove(t *testing.T) {
	testCases := []struct {
		testName           string
		initialOrientation Direction
		wantCoordinate     Coordinate
	}{
		{"up", Up, Coordinate{0, 1}},
		{"right", Right, Coordinate{1, 0}},
		{"down", Down, Coordinate{0, -1}},
		{"left", Left, Coordinate{-1, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			r := Robot{tc.initialOrientation, Coordinate{0, 0}}

			r.Move()

			if diff := cmp.Diff(tc.wantCoordinate, r.Location); diff != "" {
				t.Errorf("Coordinate difference, -want +got:\n%s", diff)
			}

		})
	}
}
