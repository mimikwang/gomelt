package gomelt

import "math"

// calcTm calculates the Tm given Ds, Dh, Dna Concentrations.  Ds should be
// corrected for salt concentration.  The units are as follows:
//
//	Tm  : C
//	Ds  : cal / mol K
//	Dh  : kcal / mol
//	Dna : nM
//
func calcTm(Ds, Dh, Dna float64, selfComp bool) float64 {
	Dna *= 1e-9
	Tm := Dh * (1000 / (Ds + r*math.Log(Dna/denom(selfComp))))
	Tm -= tKelvin
	return Tm
}

// denom retruns the appropriate denominator factor for calculating the Tm
// based on the self complementary attribute.
func denom(selfComp bool) float64 {
	if selfComp {
		return 1.0
	}
	return 4.0
}
