package calculus

// Find greatest common divisor (GCD) via Euclidean algorithm.
func FindGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
