package main

import (
	"fmt"

	"github.com/shahinam/cs-go/numerical"
)

func main() {
	// GCD Euclid's Algo.
	var m uint = 119
	var n uint = 544
	g := numerical.GcdEuclid(m, n)
	fmt.Print("GCD of ", m, " and ", n, " is ")
	fmt.Println(g)
}
