// nutritional_datatype.go
package score

import (
	"nutriscorev1/components"
)

// NutritionalData represents the source nutritional data used for the calculation
type NutritionalData struct {
	Energy              components.EnergyKJ
	Sugars              components.SugarGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}
