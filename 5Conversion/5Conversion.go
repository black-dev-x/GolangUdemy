package main

import (
	"fmt"
)

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// this does work
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
	//43.742000579833984 of type float64 is the output

}
