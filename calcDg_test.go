package gomelt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcDg(t *testing.T) {
	testCases := []struct {
		name     string
		Dh       float64
		Ds       float64
		T        float64
		expected float64
	}{
		{"Should calculate the Dg", 100.1, 3.4, 25.1, 99.08},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := calcDg(testCase.Dh, testCase.Ds, testCase.T)
			assert.InDelta(t, testCase.expected, actual, 0.01)
		})
	}
}
