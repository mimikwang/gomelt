package gomelt

// CalcDg calculates Dg when given Dh, Ds, and T.  The units are as follows:
//
//	Dg : kcal / mol
//	Dh : kcal / mol
//	Ds : cal / mol K
//	T  : C
func CalcDg(Dh, Ds, T float64) float64 {
	Dg := Dh - (Ds/1000)*(T+TKelvin)
	return Dg
}
