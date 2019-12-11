package intcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGravityComputer(t *testing.T) {
	testCases := []struct {
		initialProgram program
		finalProgram   program
	}{
		{program{1, 0, 0, 0, 99}, program{2, 0, 0, 0, 99}},
		{program{2, 3, 0, 3, 99}, program{2, 3, 0, 6, 99}},
		{program{2, 4, 4, 5, 99, 0}, program{2, 4, 4, 5, 99, 9801}},
		{program{1, 1, 1, 4, 99, 5, 6, 0, 99}, program{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tc := range testCases {

		t.Run(tc.initialProgram.String(), func(t *testing.T) {
			c := NewComputer(nil, tc.initialProgram)
			c.Run()
			if diff := cmp.Diff(tc.finalProgram, c.Program); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
