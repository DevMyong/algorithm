func solution(n int, lost []int, reserve []int) int {
	clothes := make([]int, n+2)

	//initial
	for i := 1; i < n+1; i++ {
		clothes[i] = 1
	}
	//over clothe
	for _, r := range reserve {
		clothes[r] = 2
	}
	//stolen
	for _,l := range lost {
		clothes[l]--
	}


	for i:=1; i<n+1; i++{
		if clothes[i] == 0 && clothes[i-1] == 2{
			clothes[i-1]--
			clothes[i]++
		}else if clothes[i] == 0 && clothes[i+1] == 2{
			clothes[i+1]--
			clothes[i]++
		}
	}
	
	res := 0
	for i:=1;i<n+1;i++{
		if clothes[i] == 2{
			clothes[i] = 1
		}
		res += clothes[i]
	}
	return res
}