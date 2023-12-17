// protein.go
package components

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

// ProteinGram represents the amount of protein in grams/100g
type ProteinGram float64

var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

// GetPoints returns the nutritional score
func (p ProteinGram) GetPoints(st types.ScoreType) int {
	return utilities.GetPointsFromRange(float64(p), proteinLevels)
}
