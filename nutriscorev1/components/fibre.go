// fibre.go
package components

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

// FibreGram represents the amount of fibre in grams/100g
type FibreGram float64

var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}

// GetPoints returns the nutritional score
func (f FibreGram) GetPoints(st types.ScoreType) int {
    return utilities.GetPointsFromRange(float64(f), fibreLevels)
}
