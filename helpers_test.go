package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindCombinations(t *testing.T) {
	testCases := []struct {
		target float64
		want   [][]float64
	}{
		{target: 25.0,
			want: [][]float64{{10.0, 15.0}, {25.0}},
		},
		{target: 72.5,
			want: [][]float64{{45.0, 15.0, 10.0, 2.5}, {45.0, 25.0, 2.5}},
		},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("Target Weight: %.2f", tc.target)
		t.Run(testName, func(t *testing.T) {
			r := NewResults()
			r.findCombinations(tc.target)
			got := r.Combinations
			fmt.Println(got)
			for i, want := range tc.want {
				assert.ElementsMatch(t, want, got[i])
			}
		})
	}
}
