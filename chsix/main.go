package main

import (
	"flag"
	"fmt"whatsApp
	"os"

	"github.com/emilearthur/learning-go/chsix/current"
	"github.com/emilearthur/learning-go/chsix/power"
	"github.com/emilearthur/learning-go/chsix/power/ir"
	"github.com/emilearthur/learning-go/chsix/power/vr"
	res "github.com/emilearthur/learning-go/chsix/resistor"
	"github.com/emilearthur/learning-go/chsix/volt"
)

var (
	op string
	v  float64
	r  float64
	i  float64
	p  float64

	usage = "Usage: ./circ <command> [arguments]\n" +
		"Valid command {V | Vpi | R | Rvp | I | Ivp | P | Pir | Pvr}"
)

func init() { // intialization parsing of flag values
	// flag uses os.Args to provide processing of structured cmd-line args.
	flag.Float64Var(&v, "v", 0.0, "Voltage value (volt)")
	flag.Float64Var(&r, "r", 0.0, "Resistance value (ohms)")
	flag.Float64Var(&i, "i", 0.0, "Current value (amp)")
	flag.Float64Var(&p, "p", 0.0, "Electrical power (watt)")
	flag.StringVar(&op, "op", "V", "Command - one of { V | Vpi | R | Rvp | I | Ivp | P | Pir | Pvr }")
}

// simple voltage, resistance, current calculator
// usuage ./circ -op=<operation> arguments
// example: ./circ -op=V -r 12 -i 0.5
// example: ./circ -op=I -v 15 -r 100000
func main() {
	flag.Parse() // Parse any flags provided as cmd-line.
	switch op {
	case "V", "v":
		fmt.Printf("V= %0.2f volts\n", volt.V(i, r))
	case "Vpi", "vpi":
		fmt.Printf("Vpi= %0.2f volts\n", volt.Vpi(p, i))
	case "R", "r":
		fmt.Printf("R= %0.2f Ohms\n", res.R(v, i))
	case "I", "i":
		fmt.Printf("I= %0.2f amps\n", current.I(v, r))
	case "Ivp", "ivp":
		fmt.Printf("I= %0.2f amps\n", current.Ivp(v, p))
	case "P", "p":
		fmt.Printf("I= %0.2f watts\n", power.P(v, i))
	case "Pir", "pir":
		fmt.Printf("I= %0.2f watts\n", ir.P(i, r))
	case "Pvr", "pvr":
		fmt.Printf("I= %0.2f watts\n", vr.P(v, r))
	case "Rpi", "rpi":
		fmt.Printf("I= %0.2f watts\n", res.Rpi(p, i))
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
