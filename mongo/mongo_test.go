package mongo

import (
	"reflect"
	"testing"
)

func TestFilterIncrements(t *testing.T) {
	testCases := []struct {
		increment  float64
		lowerBound float64
		upperBound float64
		expected   []float64
	}{
		{0.5, 0.0, 2.0, []float64{0.0, 0.5, 1.0, 1.5, 2.0, 2.5}},
		{0.1, 1.0, 1.5, []float64{1.0, 1.1, 1.2, 1.3, 1.4, 1.5}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := FilterIncrements(tc.increment, tc.lowerBound, tc.upperBound)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
