func solution(x int) bool {
	dup := x
	sum := 0

	for dup > 0 {
		sum += dup%10
		dup = dup/10
	}
	return x % sum == 0
}