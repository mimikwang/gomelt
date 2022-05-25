package gomelt

// calcDg calculates Dg when given Dh, Ds, and T.  The units are as follows:
//
//	Dg : kcal / mol
//	Dh : kcal / mol
//	Ds : cal / mol K
//	T  : C
//
func calcDg(Dh, Ds, T float64) float64 {
	Dg := Dh - (Ds/1000)*(T+tKelvin)
	return Dg
}
