func solution(s string) string {
	idx := 0
	res := ""
	for i, ch := range s {
		if ch == ' '{
			idx = i+1
		}
		if ch >= 'a' {
			if (i-idx)%2 == 0{
				ch -= 32
			}
		}else if ch <='Z'{
			if (i-idx)%2 == 1{
				ch += 32
			}
		}
		res += string(ch)
	}
	return res
}