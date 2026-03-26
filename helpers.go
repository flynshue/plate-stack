package main

var defaultPlateSet = []Plate{
	{Weight: 2.5, Quantity: 1},
	{Weight: 5.0, Quantity: 1},
	{Weight: 10.0, Quantity: 1},
	{Weight: 15.0, Quantity: 1},
	{Weight: 25.0, Quantity: 1},
	{Weight: 45.0, Quantity: 1},
}

type Plate struct {
	Weight   float64 `schema:"weight"`
	Quantity int     `schema:"quantity"`
}

type PlateForm struct {
	Plates       []Plate `schema:"plate"`
	PlateSet     []float64
	Barbell      float64 `schema:"barbell"`
	TargetWeight float64 `schema:"target"`
	Combinations [][]float64
}

func newPlateForm() *PlateForm {
	return &PlateForm{
		Plates:       defaultPlateSet,
		PlateSet:     []float64{},
		Combinations: make([][]float64, 0),
	}
}

func (p *PlateForm) UpdatePlateSet(plates []float64) {
	p.PlateSet = plates
}

func (p *PlateForm) buildPlateSet() {
	set := make([]float64, 0)
	for _, plate := range p.Plates {
		for i := 0; i < plate.Quantity; i++ {
			set = append(set, plate.Weight)
		}
	}
	p.PlateSet = set
}

func (p *PlateForm) findCombinations() {
	if len(p.Plates) > 0 {
		p.buildPlateSet()
	}
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
		if i > startIdx && p.PlateSet[i] == p.PlateSet[i-1] {
			continue
		}
		if p.PlateSet[i] > target {
			continue
		}
		current = append(current, p.PlateSet[i])
		newTarget := target - p.PlateSet[i]
		p.backTrack(newTarget, i+1, current)
		current = current[:len(current)-1]
	}
}
