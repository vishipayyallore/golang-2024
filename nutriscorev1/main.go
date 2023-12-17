package main

import (
	"fmt"
	"nutriscorev1/score"
)

func main() {

	fmt.Println("=========================================")
	fmt.Println("Nutritional score")
	fmt.Println("=========================================")

	ns := score.GetNutritionalScore(score.NutritionalData{
		Energy:              score.EnergyFromKcal(0),
		Sugars:              score.SugarGram(10),
		SaturatedFattyAcids: score.SaturatedFattyAcidsGram(2),
		Sodium:              score.SodiumMilligram(500),
		Fruits:              score.FruitsPercent(60),
		Fibre:               score.FibreGram(4),
		Protein:             score.ProteinGram(2),
	}, score.Food)

	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())

	fmt.Println("-----------------------------------------")
}
