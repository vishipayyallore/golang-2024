// nutritional_datatype.go
package score

import (
	"nutriscorev1/components"
)

// NutritionalData represents the source nutritional data used for the calculation
type NutritionalData struct {
	Energy              components.EnergyKJ
	Sugars              components.SugarGram
	SaturatedFattyAcids components.SaturatedFattyAcidsGram
	Sodium              components.SodiumMilligram
	Fruits              components.FruitsPercent
	Fibre               components.FibreGram
	Protein             components.ProteinGram
	IsWater             bool
}
