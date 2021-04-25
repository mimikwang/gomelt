# gomelt

[![mimikwang](https://circleci.com/gh/mimikwang/gomelt.svg?style=shield)](https://circleci.com/gh/mimikwang/gomelt)

## Overview
`gomelt` is a golang implementation of DNA oligo melting temperature calculations based on the Nearest-neighbor method.

Specific information for this can be found on this [Wikipedia page](https://en.wikipedia.org/wiki/Nucleic_acid_thermodynamics#Nearest-neighbor_method).

## Installation
```
$ go get github.com/mimikwang/gomelt
```

## Examples
Calculate melting temperature:

```go
package main

import (
	"fmt"

	"github.com/mimikwang/gomelt"
)

func main() {
	// Set up thermodynamics parameters
	parameters := gomelt.ThermoParam{
		Monovalent: 50.0,
		Divalent:   1.5,
		Dntp:       0.2,
		Dna:        200.0,
		T:          37.0,
	}

	// Calculate melting thermodynamics
	result, _ := gomelt.OligoTm("AAAGGCCTT", parameters)

	fmt.Println(result.Tm)
	// 20.24
	fmt.Println(result.Dh)
	// -59.6
	fmt.Println(result.Ds)
	// -169.74
	fmt.Println(result.Dg)
	// -6.96

}
```