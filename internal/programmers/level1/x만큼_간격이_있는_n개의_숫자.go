func solution(x int, n int) []int64 {
	arr := make([]int64, 0)
	gap := x
	for i:=0;i<n;i++{
		arr = append(arr, int64(x))
		x = x+gap
	}

	return arr
}