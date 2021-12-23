package resistor

// Rser calculates resistors in series
func Rser(resistors ...float64) (Rtotal float64) {
	for _, r := range resistors {
		Rtotal = Rtotal + r
	}
	return
}

// Rpar calculates resistors in parallel
func Rpar(resistors ...float64) (Rtotal float64) {
	for _, r := range resistors {
		Rtotal = Rtotal + recip(r) // since recip is part of the resistor package the golang complier resolves it.
	}
	return
}
