// scoring_utils.go
package utilities

// getPointsFromRange calculates points based on a value and a range
func GetPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}
