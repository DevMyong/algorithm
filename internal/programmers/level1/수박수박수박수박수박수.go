func solution(n int) string {
	su := "수"
	bak := "박"
	
	str := ""
	
	for i:=0;i<n;i++{
		if i%2 == 0{
			str+=su
		}else {
			str+=bak
		}
	}
	return str
}