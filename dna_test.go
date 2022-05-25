package gomelt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDnaBase(t *testing.T) {
	testCases := []struct {
		name     string
		base     byte
		expected bool
	}{
		{"Should return true for A", 'A', true},
		{"Should return true for T", 'T', true},
		{"Should return true for C", 'C', true},
		{"Should return true for G", 'G', true},
		{"Should return false for Z", 'Z', false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := isDnaBase(testCase.base)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
