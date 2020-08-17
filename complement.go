package gomelt

import "math"

// TabComplement maps complement bases
var TabComplement = map[byte]byte{
	'A': 'T',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

// Complement complements a sequence
func Complement(sequence string) string {
	sequenceComp := ""
	for i := 0; i < len(sequence); i++ {
		if comp, found := TabComplement[sequence[i]]; found {
			sequenceComp += string(comp)
		} else {
			// If base not found in the complement table, return original
			// base
			sequenceComp += string(sequence[i])
		}
	}

	return sequenceComp
}

// IsSelfComplement checks to see if a sequence is self complementary and
// return true if it is and false if it is not
func IsSelfComplement(sequence string) bool {
	length := len(sequence)

	// If empty string, return false
	if length < 1 {
		return false
	}

	// If odd number of bases, return false
	if math.Mod(float64(length), 2) == 1 {
		return false
	}

	midPoint := length / 2
	for i := 0; i < midPoint; i++ {
		start := sequence[i]
		end := sequence[length-i-1]
		if end != TabComplement[start] {
			return false
		}
	}

	return true
}
