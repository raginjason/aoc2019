package intcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLargeNumber(t *testing.T) {
	testCases := []struct {
		program Program
		want    int
	}{
		{Program{104, 1125899906842624, 99}, 1125899906842624},
	}

	for _, tc := range testCases {
		t.Run(tc.program.String(), func(t *testing.T) {
			in := make(chan int)
			out := make(chan int)
			c := NewComputer(in, tc.program, out)
			go c.Run()

			var got int
			if val, ok := <-out; ok {
				got = val
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
