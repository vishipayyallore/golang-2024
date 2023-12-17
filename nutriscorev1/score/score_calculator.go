package score

import (
	"github.com/vishipayyallore/learn-golang-2024/nutriscorev1/components"
)

// GetNutritionalScore calculates the nutritional score for nutritional data n of type st
func GetNutritionalScore(n components.NutritionalData, st components.ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	// Water is always graded A page 30
	if st != components.Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)

		// negative points are the negative things like calories, sugars, saturated fats, and sodium
		// positives are fruit points, fiber points, and proteins
		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)

		if st == components.Cheese {
			// Cheeses always use (negative - positive) page 29
			value = negative - positive
		} else {
			// page 27
			if negative >= 11 && fruitPoints < 5 {
				value = negative - fibrePoints - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

// GetNutriScore returns the Nutri-Score rating
func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == components.Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == components.Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}
