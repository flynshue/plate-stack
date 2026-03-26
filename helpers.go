package main

var defaultPlateSet = []Plate{
	{Weight: 2.5, Quantity: 1},
	{Weight: 5.0, Quantity: 1},
	{Weight: 10.0, Quantity: 1},
	{Weight: 15.0, Quantity: 1},
	{Weight: 25.0, Quantity: 1},
	{Weight: 35.0, Quantity: 0},
	{Weight: 45.0, Quantity: 1},
}

type Plate struct {
	Weight   float64 `schema:"weight"`
	Quantity int     `schema:"quantity"`
}

type PlateForm struct {
	Plates       []Plate `schema:"plate"`
	Barbell      float64 `schema:"barbell"`
	TargetWeight float64 `schema:"target"`
	Combinations [][]float64
}

func newPlateForm() *PlateForm {
	return &PlateForm{
		Plates:       defaultPlateSet,
		Combinations: make([][]float64, 0),
	}
}

func (p *PlateForm) buildPlateSet() []float64 {
	set := make([]float64, 0)
	for _, plate := range p.Plates {
		for i := 0; i < plate.Quantity; i++ {
			set = append(set, plate.Weight)
		}
	}
	return set
}

func (p *PlateForm) findCombinations() {
	plateSet := p.buildPlateSet()
	target := (p.TargetWeight - p.Barbell) / 2.0
	current := make([]float64, 0)
	p.backTrack(target, 0, current, plateSet)
}

func (p *PlateForm) backTrack(target float64, startIdx int, current []float64, plateSet []float64) {
	if target == 0.0 {
		temp := make([]float64, len(current))
		copy(temp, current)
		p.Combinations = append(p.Combinations, temp)
		return
	}
	for i := startIdx; i < len(plateSet); i++ {
		if i > startIdx && plateSet[i] == plateSet[i-1] {
			continue
		}
		if plateSet[i] > target {
			continue
		}
		current = append(current, plateSet[i])
		newTarget := target - plateSet[i]
		p.backTrack(newTarget, i+1, current, plateSet)
		current = current[:len(current)-1]
	}
}
