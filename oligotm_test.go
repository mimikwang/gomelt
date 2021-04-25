package gomelt_test

import (
	"testing"

	"github.com/mimikwang/gomelt"
)

func TestOligoTm(t *testing.T) {
	testCases := []struct {
		name        string
		sequence    string
		thermoParam gomelt.ThermoParam
		expected    gomelt.ThermoResult
		expectErr   bool
	}{
		{
			"Normal",
			"AAAGGCCTT",
			gomelt.ThermoParam{Monovalent: 50.0, Divalent: 1.5, Dntp: 0.2, Dna: 200, T: 37.0},
			gomelt.ThermoResult{Dh: -59.6, Ds: -169.74, Dg: -6.96, Tm: 20.24},
			false,
		},
		{
			"Unrecognized bases",
			"AAAZAA",
			gomelt.ThermoParam{},
			gomelt.ThermoResult{Dh: 0.0, Ds: 0.0, Dg: 0.0, Tm: 0.0},
			true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := gomelt.OligoTm(testCase.sequence, testCase.thermoParam)

			// Check for errors
			if testCase.expectErr && err == nil {
				t.Error()
			}
			if !testCase.expectErr && err != nil {
				t.Error()
			}

			// Check values
			if !AlmostEqual(actual.Tm, testCase.expected.Tm, 0.01) {
				t.Error()
			}
			if !AlmostEqual(actual.Dh, testCase.expected.Dh, 0.01) {
				t.Error()
			}
			if !AlmostEqual(actual.Ds, testCase.expected.Ds, 0.01) {
				t.Error()
			}
			if !AlmostEqual(actual.Dg, testCase.expected.Dg, 0.01) {
				t.Error()
			}
		})
	}
}
