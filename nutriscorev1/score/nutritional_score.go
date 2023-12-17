package score

import "github.com/vishipayyallore/learn-golang-2024/nutriscorev1/components"

// NutritionalScore represents the nutritional score information
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType components.ScoreType
}
