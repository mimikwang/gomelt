package gomelt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaltCorrDs(t *testing.T) {
	testCases := []struct {
		name       string
		ds         float64
		monovalent float64
		divalent   float64
		dntp       float64
		length     int
		expected   float64
		expectErr  bool
	}{
		{
			"Should correct for salt concentration",
			100.0,
			10.0,
			10.0,
			10.0,
			20.0,
			67.80,
			false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := saltCorrDs(
				testCase.ds,
				testCase.monovalent,
				testCase.divalent,
				testCase.dntp,
				testCase.length,
			)
			assert.InDelta(t, testCase.expected, actual, 0.01)
			if testCase.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDivToMono(t *testing.T) {
	testCases := []struct {
		name      string
		divalent  float64
		dntp      float64
		expected  float64
		expectErr bool
	}{
		{
			"Should convert divalent salt conc to mono salt conc",
			10.0,
			5.0,
			268.33,
			false,
		},
		{
			"Should return error if divalent is negative",
			-10.0,
			10.0,
			0.0,
			true,
		},
		{
			"Should return error if dntp is negative",
			10.0,
			-10.0,
			0.0,
			true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := divToMono(testCase.divalent, testCase.dntp)
			assert.InDelta(t, testCase.expected, actual, 0.01)
			if testCase.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
