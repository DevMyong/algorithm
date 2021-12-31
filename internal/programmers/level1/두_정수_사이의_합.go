func solution(a int, b int) int64 {
	var res int64
	if a > b{
		tmp := a
		a = b
		b = tmp
	}
	for i:=a;i<=b;i++{
        res += int64(i)
	}
	
	return res
}