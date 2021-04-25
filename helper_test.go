package gomelt_test

import "math"

// AlmostEqual is a helper function that compares two floats and check that
// they're within the provided margin.  Returns `true` if it does and `false`
// if it does not.
func AlmostEqual(value1, value2, margin float64) bool {
	diff := math.Abs(value1 - value2)
	if diff > margin {
		return false
	}
	return true
}
