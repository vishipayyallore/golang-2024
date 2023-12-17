// scoring_utils_test.go
package utilities

import "testing"

func TestGetPointsFromRange(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		steps    []float64
		expected int
	}{
		{"Value in first step", 2.5, []float64{5, 10, 15}, 0},
		{"Value in middle step", 12.5, []float64{5, 10, 15}, 3},
		{"Value in last step", 17.5, []float64{5, 10, 15}, 3},
		{"Value below all steps", 3, []float64{5, 10, 15}, 0},
		{"Value above all steps", 20, []float64{5, 10, 15}, 3},
		{"Empty steps", 10, []float64{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPointsFromRange(tt.value, tt.steps)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
