package gomelt

import "math"

// CalcTm calculates the Tm given Ds, Dh, Dna Concentrations.  Ds should be
// corrected for salt concentration.  The units are as follows:
//
//	Tm  : C
//	Ds  : cal / mol K
//	Dh  : kcal / mol
//	Dna : nM
func CalcTm(Ds, Dh, Dna float64, selfComp bool) float64 {
	denom := 4.0
	if selfComp {
		denom = 1.0
	}
	Dna *= 1e-9
	Tm := Dh * (1000 / (Ds + R*math.Log(Dna/denom)))
	Tm -= TKelvin

	return Tm
}
