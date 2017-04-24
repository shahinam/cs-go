package numerical

import "testing"

type testCases struct {
	m   uint
	n   uint
	gcd uint
}

var tests = []testCases{
	{119, 544, 17},
	{100, 0, 100},
	{30315475, 24440870, 31415},
}

func TestGcdEuclid(t *testing.T) {
	for _, test := range tests {
		g := GcdEuclid(test.m, test.n)
		if g != test.gcd {
			t.Error(
				"For", test.m, "and", test.n,
				"expected", test.gcd,
				"got", g,
			)
		}
	}
}
