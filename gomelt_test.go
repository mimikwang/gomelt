package gomelt_test

import (
	"testing"

	"github.com/mimikwang/gomelt"
	"github.com/stretchr/testify/assert"
)

func TestMeltingTemp(t *testing.T) {
	testCases := []struct {
		name      string
		sequence  string
		params    gomelt.Params
		expected  gomelt.Result
		expectErr bool
	}{
		{
			"Normal",
			"AAAGGCCTT",
			gomelt.Params{Monovalent: 50.0, Divalent: 1.5, Dntp: 0.2, Dna: 200, T: 37.0},
			gomelt.Result{Dh: -59.6, Ds: -169.74, Dg: -6.96, Tm: 20.24},
			false,
		},
		{
			"Unrecognized bases",
			"AAAZAA",
			gomelt.Params{},
			gomelt.Result{Dh: 0.0, Ds: 0.0, Dg: 0.0, Tm: 0.0},
			true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := gomelt.MeltingTemp(testCase.sequence, testCase.params)

			// Check errors
			if testCase.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			// Check numerical values
			assert.InDelta(t, testCase.expected.Tm, actual.Tm, 0.01)
			assert.InDelta(t, testCase.expected.Dh, actual.Dh, 0.01)
			assert.InDelta(t, testCase.expected.Ds, actual.Ds, 0.01)
			assert.InDelta(t, testCase.expected.Dg, actual.Dg, 0.01)
		})
	}
}
