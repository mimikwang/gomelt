package gomelt

// ThermoResult holds thermodynamics result
type ThermoResult struct {
	Dh float64 // kcal / mol
	Ds float64 // cal / mol K
	Dg float64 // kcal / mol
	Tm float64 // C
}

var oligoTmErr ThermoResult = ThermoResult{0, 0, 0, 0}

// ThermoParam holds thermodynamics parameters
type ThermoParam struct {
	Monovalent float64 // mM
	Divalent   float64 // mM
	Dntp       float64 // mM
	Dna        float64 // nM
	T          float64 // C
}

// OligoTm calculates the oligo's melting temperature given a sequence and
// thermodynamic parameters.  It returns a ThermoResult with Dh, Ds, Dg, and
// Tm.
func OligoTm(sequence string, param ThermoParam) (ThermoResult, error) {
	dh := HInit
	ds := SInit
	length := len(sequence)

	// Symmetry Penalty
	selfComp := IsSelfComplement(sequence)
	if selfComp {
		dh += HSymmetry
		ds += SSymmetry
	}

	// Terminal A/T Penalties
	if sequence[0] == 'A' || sequence[0] == 'T' {
		dh += HTerminalAT
		ds += STerminalAT
	}
	if sequence[length-1] == 'A' || sequence[length-1] == 'T' {
		dh += HTerminalAT
		ds += STerminalAT
	}

	// Add Internal
	for i := 0; i < length-1; i++ {
		subSeq := sequence[i : i+2]
		if _, found := TabThermo[subSeq]; !found {
			return oligoTmErr, ErrInvalidBase
		}
		dh += TabThermo[subSeq].dh
		ds += TabThermo[subSeq].ds
	}

	// Salt Correct Ds
	ds, err := SaltCorrDs(ds, param.Monovalent, param.Divalent, param.Dntp, length)
	if err != nil {
		return oligoTmErr, err
	}

	// Calc Tm
	tm := CalcTm(ds, dh, param.Dna, selfComp)

	// Calc Dg
	dg := CalcDg(dh, ds, param.T)

	// Return
	return ThermoResult{Dh: dh, Dg: dg, Ds: ds, Tm: tm}, nil
}
