import "sort"

func solution(n int64) (res int) {
	num := make([]int, 0)
	for n > 0 {
		num = append(num,int(n%10))
		n /= 10
	}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))

	for _,m := range num{
		res = res * 10 + m
	}
	return
}