package main

import (
	"fmt"
	"nutriscorev1/components"
	"nutriscorev1/score"
	"nutriscorev1/types"
)

func main() {

	fmt.Println("=========================================")
	fmt.Println("Nutritional score")
	fmt.Println("=========================================")

	ns := score.GetNutritionalScore(score.NutritionalData{
		Energy:              components.EnergyFromKcal(0),
		Sugars:              components.SugarGram(10),
		SaturatedFattyAcids: components.SaturatedFattyAcidsGram(2),
		Sodium:              components.SodiumMilligram(500),
		Fruits:              components.FruitsPercent(60),
		Fibre:               components.FibreGram(4),
		Protein:             components.ProteinGram(2),
	}, types.Food)

	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())

	fmt.Println("-----------------------------------------")
}
