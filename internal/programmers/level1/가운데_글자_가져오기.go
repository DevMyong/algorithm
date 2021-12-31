func solution(s string) string {
	if len(s)%2 == 1{
		return string(s[len(s)/2])
	}else {
		return s[len(s)/2-1:len(s)/2+1]
	}
}