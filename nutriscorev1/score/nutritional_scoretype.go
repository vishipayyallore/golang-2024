// nutritional_score.go
package score

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

// NutritionalScore represents the nutritional score
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType types.ScoreType
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

// GetNutriScore returns the Nutri-Score rating
func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == types.Food {
		return scoreToLetter[utilities.GetPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == types.Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[utilities.GetPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}
