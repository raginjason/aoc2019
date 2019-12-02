package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strconv"
	"strings"
	"testing"
)

func TestGravityComputer(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tc := range testCases {
		var inputStrs []string
		for _, num := range tc.input {
			s := strconv.Itoa(num)
			inputStrs = append(inputStrs, s)
		}
		inputName := strings.Join(inputStrs, ",")

		var wantStrs []string
		for _, num := range tc.want {
			s := strconv.Itoa(num)
			wantStrs = append(wantStrs, s)
		}
		wantName := strings.Join(wantStrs, ",")

		t.Run(fmt.Sprintf("%s->%s", inputName, wantName), func(t *testing.T) {
			got := GravityComputer(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}
