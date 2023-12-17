// energy.go
package components

import (
	"nutriscorev1/types"
	"nutriscorev1/utilities"
)

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}

// EnergyKJ represents the energy density in kJ/100g
type EnergyKJ float64

// EnergyFromKcal converts energy density from kcal to EnergyKJ
func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

// GetPoints returns the nutritional score
func (e EnergyKJ) GetPoints(st types.ScoreType) int {
	if st == types.Beverage {
		return utilities.GetPointsFromRange(float64(e), energyLevelsBeverage)
	}
	return utilities.GetPointsFromRange(float64(e), energyLevels)
}
