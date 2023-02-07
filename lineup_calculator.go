package main

type calculator interface {
	calculate(contents []content, interactions []interaction) *map[int]int
}

func newCalculator() calculator {
	return &basicCalculator{}
}

type basicCalculator struct {
}

func (lc *basicCalculator) calculate(contents []content, interactions []interaction) *map[int]int {
	lineup := map[int]int{}

	for _, interaction := range interactions {
		lineup[interaction.contentId] += 1
	}

	return &lineup
}
