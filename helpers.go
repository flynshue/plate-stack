package main

var defaultPlateSet = []float64{2.5, 5.0, 10.0, 15.0, 25.0, 45.0}

type PlateForm struct {
	PlateSet     []float64
	Barbell      float64 `schema:"barbell"`
	TargetWeight float64 `schema:"target"`
	Combinations [][]float64
}

func newPlateForm() *PlateForm {
	return &PlateForm{
		PlateSet:     defaultPlateSet,
		Combinations: make([][]float64, 0),
	}
}

func (p *PlateForm) UpdatePlateSet(plates []float64) {
	p.PlateSet = plates
}

func (p *PlateForm) findCombinations() {
	target := (p.TargetWeight - p.Barbell) / 2.0
	current := make([]float64, 0)
	p.backTrack(target, 0, current)
}

func (p *PlateForm) backTrack(target float64, startIdx int, current []float64) {
	if target == 0.0 {
		temp := make([]float64, len(current))
		copy(temp, current)
		p.Combinations = append(p.Combinations, temp)
		return
	}
	for i := startIdx; i < len(p.PlateSet); i++ {
		if p.PlateSet[i] > target {
			continue
		}
		current = append(current, p.PlateSet[i])
		newTarget := target - p.PlateSet[i]
		p.backTrack(newTarget, i+1, current)
		current = current[:len(current)-1]
	}
}
