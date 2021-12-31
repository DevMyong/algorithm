func solution(s string, n int) string {
	r := make([]rune, len(s))
	
	for i, ch := range s {
		if ch <= 'z' && ch >= 'a' {
			r[i] = (ch + rune(n) - 'a')%26 + 'a'
		}else if ch <= 'Z' && ch >= 'A' {
			r[i] = (ch + rune(n) - 'A')%26 + 'A'
		}else {
			r[i] = ch
		}
	}
	return string(r)
}