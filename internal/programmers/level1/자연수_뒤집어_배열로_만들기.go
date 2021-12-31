func solution(n int64) []int64 {
	res := make([]int64, 0)

	for n>0 {
		res = append(res, n%10)
		n /= 10
	}
	return res
}