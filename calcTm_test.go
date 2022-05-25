package gomelt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcTm(t *testing.T) {
	testCases := []struct {
		name     string
		Ds       float64
		Dh       float64
		Dna      float64
		selfComp bool
		expected float64
	}{
		{
			"Should calculate the Tm",
			1000.2,
			400.1,
			100,
			false,
			141.28,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := calcTm(
				testCase.Ds, testCase.Dh, testCase.Dna, testCase.selfComp)
			assert.InDelta(t, testCase.expected, actual, 0.01)
		})
	}
}

func TestDenom(t *testing.T) {
	testCases := []struct {
		name     string
		selfComp bool
		expected float64
	}{
		{"Should return 4.0 if not self complementary", false, 4.0},
		{"Should return 1.0 if self complementary", true, 1.0},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := denom(testCase.selfComp)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
