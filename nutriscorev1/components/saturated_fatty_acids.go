// saturated_fatty_acids.go
package components

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

// SaturatedFattyAcidsGram represents the amount of saturated fatty acids in grams/100g
type SaturatedFattyAcidsGram float64

var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

// GetPoints returns the nutritional score
func (s SaturatedFattyAcidsGram) GetPoints(st types.ScoreType) int {
	return utilities.GetPointsFromRange(float64(s), saturatedFattyAcidsLevels)
}
