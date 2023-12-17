package score

import (
	"nutriscorev1/components/energy"
)

// GetNutritionalScore calculates the nutritional score for nutritional data n of type st
func GetNutritionalScore(n energy.NutritionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	// Just an example for Energy, you would extend this for other components
	negative = n.Energy.GetPoints(st)

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

// GetNutriScore returns the Nutri-Score rating
func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}
