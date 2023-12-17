// sodium.go
package components

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

// SodiumMilligram represents the amount of sodium in mg/100g
type SodiumMilligram float64

var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}

// GetPoints returns the nutritional score
func (s SodiumMilligram) GetPoints(st types.ScoreType) int {
	return utilities.GetPointsFromRange(float64(s), sodiumLevels)
}

// SodiumFromSalt converts salt mg/100g content to sodium content
func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}
