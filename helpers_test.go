package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindCombinations(t *testing.T) {
	r := NewResults()
	r.findCombinations(25.0)
	want := [][]float64{{25.0}, {10.0, 15.0}}
	fmt.Println(r.Combinations)
	assert.ElementsMatch(t, want, r.Combinations)
}
