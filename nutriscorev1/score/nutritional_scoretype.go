// nutritional_score.go
package score

import "nutriscorev1/types"

// NutritionalScore represents the nutritional score
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType types.ScoreType
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}
