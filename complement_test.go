package gomelt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplementBase(t *testing.T) {
	testCases := []struct {
		name     string
		base     byte
		expected byte
	}{
		{"A should complement to T", 'A', 'T'},
		{"T should complement to A", 'T', 'A'},
		{"C should complement to G", 'C', 'G'},
		{"G should complement to C", 'G', 'C'},
		{"Non DNA base should return itself", 'Z', 'Z'},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := complementBase(testCase.base)
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
			actual := isSelfComplement(testCase.sequence)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestIsOdd(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		expected bool
	}{
		{"Should return true for 1", 1, true},
		{"Should return false for 4", 4, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := isOdd(testCase.value)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
