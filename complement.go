package gomelt

import "math"

// complementBase complements a base.  The original base is returned if the
// base is not a DNA base.
func complementBase(base byte) byte {
	switch base {
	case 'A':
		return 'T'
	case 'T':
		return 'A'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	default:
		return base
	}
}

// isSelfComplement checks to see if a sequence is self complementary and
// return true if it is and false if it is not.
func isSelfComplement(sequence string) bool {
	length := len(sequence)

	// If empty string or odd number of bases, return false
	if length < 1 || isOdd(length) {
		return false
	}

	midPoint := length / 2
	for i := 0; i < midPoint; i++ {
		start := sequence[i]
		end := sequence[length-i-1]
		if end != complementBase(start) ||
			!isDnaBase(start) ||
			!isDnaBase(end) {
			return false
		}
	}
	return true
}

// isOdd returns true if the value is odd and false otherwise.
func isOdd(value int) bool {
	return math.Mod(float64(value), 2) == 1
}
