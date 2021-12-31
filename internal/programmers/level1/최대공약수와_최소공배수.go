func solution(n int, m int) []int {
	return []int{ gcd(n,m), lcm(n,m) }
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	result := a * b / gcd(a, b)
	return result
}