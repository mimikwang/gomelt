package gomelt

// Numerical constants
const (
	TKelvin     = 273.15 // Kelvin
	R           = 1.987  // Ideal gas constant in (cal / mol K)
	HInit       = 0.2    // Initial Enthalpy in (kcal / mol)
	SInit       = -5.7   // Initial Entropy in (cal / mol K)
	HSymmetry   = 0      // Symmetry Enthalpy penalty in (kcal / mol)
	SSymmetry   = -1.4   // Symmetry Entropy penalty in (cal / mol K)
	HTerminalAT = 2.2    // Terminal AT Enthalpy penalty in (kcal / mol)
	STerminalAT = 6.9    // Terminal AT Entropy penalty in (cal / mol K)
)

type thermoParam struct {
	dh float64 // kcal / mol
	ds float64 // cal / mol K
}

// TabThermo holds thermodynamics parameters
var TabThermo map[string]thermoParam = map[string]thermoParam{
	"GA": {-8.2, -22.2},
	"GG": {-8, -19.9},
	"GT": {-8.4, -22.4},
	"GC": {-9.8, -27.2},
	"CA": {-8.5, -22.7},
	"CC": {-8, -19.9},
	"CG": {-10.6, -27.2},
	"CT": {-7.8, -21},
	"AA": {-7.6, -21.3},
	"AC": {-8.4, -22.4},
	"AG": {-7.8, -21},
	"AT": {-7.2, -20.4},
	"TA": {-7.2, -21.3},
	"TC": {-8.2, -22.2},
	"TG": {-8.5, -22.7},
	"TT": {-7.6, -21.3},
}
