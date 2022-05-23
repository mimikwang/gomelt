package gomelt_test

import (
	"testing"

	"github.com/go-playground/assert"
	"github.com/mimikwang/gomelt"
)

func TestComplement(t *testing.T) {
	testCases := []struct {
		name     string
		sequence string
		expected string
	}{
		{"Normal", "AGCTTCGA", "TCGAAGCT"},
		{"Empty", "", ""},
		{"Unrecognized bases", "AZZZ", "TZZZ"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := gomelt.Complement(testCase.sequence)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestIsSelfComplement(t *testing.T) {
	testCases := []struct {
		name     string
		sequence string
		expected bool
	}{
		{"Normal", "AGCT", true},
		{"Odd number bases", "AAATT", false},
		{"Empty", "", false},
		{"Unrecognized bases", "ZZZZ", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := gomelt.IsSelfComplement(testCase.sequence)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
