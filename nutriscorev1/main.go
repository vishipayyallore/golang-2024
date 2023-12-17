package main

import (
	"fmt"
)

func main() {
	fmt.Println("Nutritional score")

	ns := score.GetNutritionalScore(energy.NutritionalData{
		Energy: score.EnergyFromKcal(0),
	}, score.Food)

	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())

	// ns := score.GetNutritionalScore(components.NutritionalData{
	// 	Energy:              components.EnergyFromKcal(0),
	// 	Sugars:              components.SugarGram(10),
	// 	SaturatedFattyAcids: components.SaturatedFattyAcidsGram(2),
	// 	Sodium:              components.SodiumMilligram(500),
	// 	Fruits:              components.FruitsPercent(60),
	// 	Fibre:               components.FibreGram(4),
	// 	Protein:             components.ProteinGram(2),
	// }, components.Food)

}
