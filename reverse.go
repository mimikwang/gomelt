package gomelt

// Reverse reverses a sequence
func Reverse(sequence string) string {
	sequenceReverse := ""
	for i := len(sequence); i > 0; i-- {
		sequenceReverse += sequence[i-1 : i]
	}

	return sequenceReverse
}
