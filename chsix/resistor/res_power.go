package resistor

// Rpi calculate Resistivity current
var Rpi func(float64, float64) float64

func init() { // init used to encapsulate custom intialization logic that when invoked when
	// packages is imported
	Rpi = func(p, i float64) float64 {
		return p / (i * i)

	}
}

// Rvp calculate Resistivity power
func Rvp(v, p float64) float64 {
	return (v * v) / p
}
