package euclid

func gcd(p, q uint) uint {
	// p vs. q shouldn't matter
	if q == 0 {
		return p
	}

	// p < q => r = p
	r := p % q

	// if r = p then p and q have swapped
	return gcd(q, r)
}
