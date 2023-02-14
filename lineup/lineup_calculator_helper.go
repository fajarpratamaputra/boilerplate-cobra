package lineup

func ToLineup(calculateResult *map[int]float64) []int {
	var lineup []int
	for i, _ := range *calculateResult {
		lineup = append(lineup, i)
	}

	return lineup
}
