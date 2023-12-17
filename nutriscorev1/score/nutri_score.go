package score

import (
	"nutriscorev1/components"
	"nutriscorev1/components/constants"
)

// NutriScore represents the Nutri-Score
type NutriScore struct {
	Value     int
	ScoreType components.ScoreType
}

// GetNutriScore returns the Nutri-Score rating
func (ns NutriScore) GetNutriScore() string {
	return constants.ScoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}
