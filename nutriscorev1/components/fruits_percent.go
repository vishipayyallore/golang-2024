// fruits_percent.go
package components

import (
	"nutriscorev1/types"
)

// FruitsPercent represents fruits, vegetables, pulses, nuts, and rapeseed, walnut, and olive oils
// as a percentage of the total
type FruitsPercent float64

// GetPoints returns the nutritional score
func (f FruitsPercent) GetPoints(st types.ScoreType) int {
    if st == types.Beverage {
        if f > 80 {
            return 10
        } else if f > 60 {
            return 4
        } else if f > 40 {
            return 2
        }
        return 0
    }
    if f > 80 {
        return 5
    } else if f > 60 {
        return 2
    } else if f > 40 {
        return 1
    }
    return 0
}
