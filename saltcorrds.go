package gomelt

import (
	"fmt"
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
	if divalent < 0 || dntp < 0 {
		err := fmt.Errorf("Divalent and Dntp should be greater or equal to 0")
		return 0.0, err
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
