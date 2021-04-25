package gomelt_test

import (
	"testing"

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
			if actual != testCase.expected {
				t.Error()
			}
		})
	}
}
