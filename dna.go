package gomelt

// isDnaBase returns true if the base is a Dna base and false otherwise.
func isDnaBase(base byte) bool {
	switch base {
	case 'A', 'C', 'G', 'T':
		return true
	}
	return false
}
