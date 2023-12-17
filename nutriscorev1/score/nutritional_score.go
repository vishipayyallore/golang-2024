package score

// NutritionalScore represents the nutritional score information
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType components.ScoreType
}
