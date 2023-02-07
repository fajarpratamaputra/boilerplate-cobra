package main

const ViewScale = 1.0
const LoveScale = 2.0
const CommentScale = 3.0
const ShareScale = 4.0

type calculator interface {
	calculate(contents []content, interactions []interaction) *map[int]float64
}

func newCalculator() calculator {
	return &basicCalculator{}
}

type basicCalculator struct {
}

func (lc *basicCalculator) calculate(contents []content, interactions []interaction) *map[int]float64 {
	lineup := map[int]float64{}

	var score float64

	for _, interaction := range interactions {
		switch interaction.behaviorType {
		case "view":
			score += ViewScale
			break
		case "like":
			score += LoveScale
			break
		case "comment":
			score += CommentScale
			break
		case "share":
			score += ShareScale
			break
		}

		lineup[interaction.contentId] += score
	}

	return &lineup
}
