package calculator

import "sort"

func SortByValueDescending(calculateResult *map[int]float64) []int {
	var keys []int
	for key := range *calculateResult {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return (*calculateResult)[keys[i]] > (*calculateResult)[keys[j]]
	})

	return keys
}

func ToLineup(calculateResult *map[int]float64) []int {
	lineup := SortByValueDescending(calculateResult)

	return lineup
}
