import "sort"

func solution(arr []int, divisor int) []int {
	var res []int
	for _, a := range arr{
		if a % divisor == 0 {
			res = append(res, a)
		}
	}
	sort.Ints(res)

	if res == nil{
		return []int{-1}
	}
	return res
}