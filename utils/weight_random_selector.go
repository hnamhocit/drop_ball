package utils

import "math/rand"

func WeightedRandomSelector(weights map[int]int) int {
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}

	randomNumber := rand.Intn(totalWeight)

	cumulativeWeight := 0
	for item, weight := range weights {
		cumulativeWeight += weight
		if randomNumber < cumulativeWeight {
			return item
		}
	}

	return 0
}
