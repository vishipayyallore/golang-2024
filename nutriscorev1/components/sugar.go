// sugar.go
package components

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

// SugarGram represents the amount of sugars in grams/100g
type SugarGram float64

var sugarsLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}

var sugarsLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

// GetPoints returns the nutritional score
func (s SugarGram) GetPoints(st types.ScoreType) int {
	if st == types.Beverage {
		return utilities.GetPointsFromRange(float64(s), sugarsLevelsBeverage)
	}
	return utilities.GetPointsFromRange(float64(s), sugarsLevels)
}
