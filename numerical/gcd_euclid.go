package numerical

// GcdEuclid calculates GCD by Euclid's algorithm.
func GcdEuclid(m uint, n uint) uint {
	if n == 0 {
		return m
	}

	r := m % n
	return GcdEuclid(n, r)
}
