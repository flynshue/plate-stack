package main

var defaultPlateSet = []float64{2.5, 5.0, 10.0, 15.0, 25.0, 45.0}

type Results struct {
	PlateSet     []float64
	Total        float64
	Combinations [][]float64
}

func NewResults() Results {
	return Results{PlateSet: defaultPlateSet, Combinations: make([][]float64, 0)}
}

func (r *Results) findCombinations(target float64) {
	current := make([]float64, 0)
	r.backTrack(target, 0, current)
}

func (r *Results) backTrack(target float64, startIdx int, current []float64) {
	if target == 0.0 {
		r.Combinations = append(r.Combinations, current)
		return
	}
	for i := startIdx; i < len(r.PlateSet); i++ {
		if r.PlateSet[i] > target {
			continue
		}
		current = append(current, r.PlateSet[i])
		newTarget := target - r.PlateSet[i]
		r.backTrack(newTarget, i+1, current)
		current = current[:len(current)-1]
	}
}
