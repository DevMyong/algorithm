import "math"

func solution(arr []int) []int {
	if len(arr) == 1{
		return []int{-1}
	}
	min := math.MaxInt32
	idx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			idx = i
		}
	}
	return append(arr[0:idx], arr[idx+1:]...)
}