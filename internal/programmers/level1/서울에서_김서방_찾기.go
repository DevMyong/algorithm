import "strconv"

func solution(seoul []string) string {
	res := "김서방은 "
	for i, kim := range seoul {
		if kim == "Kim" {
			return res+ strconv.Itoa(i) + "에 있다"
		}
	}
	return ""
}