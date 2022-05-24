package gomelt

import (
	"math"
)

// SaltCorrDs calculates the salt corrected Ds.  The units are as follows:
//
//	dsCorr     : cal / mol K
//	ds         : cal / mol K
//	monovalent : mM
//	divalent   : mM
//	dntp       : mM
func SaltCorrDs(ds, monovalent, divalent, dntp float64, length int,
) (float64, error) {
	divToMono, err := divToMono(divalent, dntp)
	if err != nil {
		return 0, err
	}

	monovalent += divToMono
	dsCorr := ds + 0.368*(float64(length)-1)*math.Log(monovalent/1000)

	return dsCorr, nil
}

// divToMono converts divalent salt concentration to monovalent salt
// concentration
func divToMono(divalent, dntp float64) (float64, error) {
	if divalent < 0 {
		return 0.0, ErrInvalidDivalent
	}
	if dntp < 0 {
		return 0.0, ErrInvalidDntp
	}

	if divalent == 0 {
		dntp = 0
	}
	if divalent < dntp {
		divalent = dntp
	}

	monovalent := 120 * math.Sqrt(divalent-dntp)
	return monovalent, nil
}
