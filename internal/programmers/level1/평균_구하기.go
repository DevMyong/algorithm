func solution(arr []int) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += float64(v)
	}
	return sum /float64(len(arr))
}