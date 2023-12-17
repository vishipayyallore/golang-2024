package components

// EnergyFromKcal converts energy density from kcal to score.EnergyKJ
func EnergyFromKcal(kcal float64) score.EnergyKJ {
	return score.EnergyKJ(kcal * 4.184)
}

// GetPoints returns the nutritional score points for energy
func (e score.EnergyKJ) GetPoints(st score.ScoreType) int {
	if st == score.Beverage {
		return getPointsFromRange(float64(e), score.EnergyLevelsBeverage)
	}
	return getPointsFromRange(float64(e), score.EnergyLevels)
}
