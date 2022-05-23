package gomelt_test

import (
	"testing"

	"github.com/go-playground/assert"
	"github.com/mimikwang/gomelt"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		sequence string
		expected string
	}{
		{"Normal", "ACGTACGT", "TGCATGCA"},
		{"Empty", "", ""},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := gomelt.Reverse(testCase.sequence)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
