import "sort"

func solution(s string) string {
	chs := make([]rune, 0)

	for _, ch := range s {
		chs = append(chs, ch)
	}
	sort.Slice(chs, func(i, j int) bool { return chs[i] > chs[j] })
	
	str :=""
	for _, ch := range chs {
		str += string(ch)
	}

	return str
}