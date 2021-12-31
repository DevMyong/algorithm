func solution(s string) bool {
	if !(len(s) == 4 || len(s) == 6) {
		return false
	}

	for _, ch := range s {
		if !(ch >= 48 && ch<=57) {
			return false
		}
	}
	return true
}