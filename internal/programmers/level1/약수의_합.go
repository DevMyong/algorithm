func solution(n int) (sum int) {
	divisor := make([]int, 0)
	keys := make(map[int]bool)

	for i:=1; i<=n; i++{

		for j:=i; j<=n; j++{
			if i*j == n {
				divisor = append(divisor, i, j)
			}
		}
	}

	res := make([]int,0)
	for _,value := range divisor{
		if _, saveValue := keys[value]; !saveValue{
			keys[value] = true
			res = append(res, value)
		}
	}
	for _,value := range res {
		sum += value
	}
	return
}