package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindCombinations(t *testing.T) {
	testCases := []struct{ target float64 }{
		{target: 95.0},
		{target: 190.0},
		{target: 195.0},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("Target Weight: %.2f", tc.target)
		t.Run(testName, func(t *testing.T) {
			form := newPlateForm()
			form.Barbell = 45.0
			form.TargetWeight = tc.target
			form.findCombinations()
			fmt.Println("Found weight combinations: ", form.Combinations)
			for _, plateSet := range form.Combinations {
				fmt.Printf("Testing combination %v\n", plateSet)
				sum := form.Barbell
				for _, weight := range plateSet {
					sum += (weight * 2)
				}
				fmt.Printf("got: %.2f\nwanted: %.2f\n", sum, tc.target)
				assert.EqualValues(t, tc.target, sum)
			}
		})
	}
}
