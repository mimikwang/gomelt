package gomelt

// Result holds thermodynamics result
type Result struct {
	Dh float64 // kcal / mol
	Ds float64 // cal / mol K
	Dg float64 // kcal / mol
	Tm float64 // C
}

// Params holds thermodynamics parameters
type Params struct {
	Monovalent float64 // mM
	Divalent   float64 // mM
	Dntp       float64 // mM
	Dna        float64 // nM
	T          float64 // C
}

// MeltingTemp calculates the oligo's melting temperature given a sequence and
// thermodynamic parameters.  It returns a ThermoResult with Dh, Ds, Dg, and
// Tm.
func MeltingTemp(sequence string, params Params) (Result, error) {
	dh := hInit
	ds := sInit
	length := len(sequence)

	// Symmetry Penalty
	selfComp := isSelfComplement(sequence)
	if selfComp {
		dh += hSymmetry
		ds += sSymmetry
	}

	// Terminal A/T Penalties
	if sequence[0] == 'A' || sequence[0] == 'T' {
		dh += hTerminalAT
		ds += sTerminalAT
	}
	if sequence[length-1] == 'A' || sequence[length-1] == 'T' {
		dh += hTerminalAT
		ds += sTerminalAT
	}

	// Add Internal
	for i := 0; i < length-1; i++ {
		subSeq := sequence[i : i+2]
		if _, found := tabThermo[subSeq]; !found {
			return Result{}, ErrInvalidBase
		}
		dh += tabThermo[subSeq].dh
		ds += tabThermo[subSeq].ds
	}

	// Salt Correct Ds
	ds, err := saltCorrDs(ds, params.Monovalent, params.Divalent, params.Dntp, length)
	if err != nil {
		return Result{}, err
	}

	// Calc Tm
	tm := calcTm(ds, dh, params.Dna, selfComp)

	// Calc Dg
	dg := calcDg(dh, ds, params.T)

	// Return
	return Result{Dh: dh, Dg: dg, Ds: ds, Tm: tm}, nil
}
