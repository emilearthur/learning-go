package volt

// V return product of current(i) and resistance(r)
func V(i, r float64) float64 {
	return i * r
}

// Vser calculates voltage in series.ls
func Vser(volts ...float64) (Vtotal float64) {
	for _, v := range volts {
		Vtotal = Vtotal + v
	}
	return
}

//Vpi returns ratio of  power and current
func Vpi(p, i float64) float64 {
	return p / i
}
