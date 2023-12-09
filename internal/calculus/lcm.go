package calculus

// Find Least Common Multiple (LCM) via Greatest Common Divisor (GCD).
func FindLCM(values ...int) int {
	result := values[0] * values[1] / FindGCD(values[0], values[1])

	for i := 2; i < len(values); i++ {
		result = FindLCM(result, values[i])
	}

	return result
}
