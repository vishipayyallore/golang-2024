// More about-score: https://en.wikipedia.org/wiki/Nutri-Score
package score

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}

// SaturatedFattyAcidsGram represents amount of saturated fatty acids in grams/100g
type SaturatedFattyAcidsGram float64

// SodiumMilligram represents amount of sodium in mg/100g
type SodiumMilligram float64

// FruitsPercent represents fruits, vegetables, pulses, nuts, and rapeseed, walnut and olive oils
// as percentage of the total
type FruitsPercent float64

// FibreGram represents amount of fibre in grams/100g
type FibreGram float64

// SodiumFromSalt converts salt mg/100g content to sodium content
func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

// GetPoints returns the nutritional score
func (sfa SaturatedFattyAcidsGram) GetPoints(st types.ScoreType) int {
	return utilities.GetPointsFromRange(float64(sfa), saturatedFattyAcidsLevels)
}

// GetPoints returns the nutritional score
func (s SodiumMilligram) GetPoints(st types.ScoreType) int {
	return utilities.GetPointsFromRange(float64(s), sodiumLevels)
}

// GetPoints returns the nutritional score
func (f FruitsPercent) GetPoints(st types.ScoreType) int {
	if st == types.Beverage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		}
		return 0
	}
	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

// GetPoints returns the nutritional score
func (f FibreGram) GetPoints(st types.ScoreType) int {
	return utilities.GetPointsFromRange(float64(f), fibreLevels)
}

// GetNutritionalScore calculates the nutritional score for nutritional data n of type st
func GetNutritionalScore(n NutritionalData, st types.ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0
	// Water is always graded A page 30
	if st != types.Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)
		//negative points are the negative things like calories (it says energy but these are what people are avoiding as these are calories)
		//sugars, saturated fats and sodium
		//positives are fruit points, fiber points and proteins
		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)

		if st == types.Cheese {
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
