func solution(num int) int {
	chance := 0
	for num != 1 && chance < 500 {
		if num % 2 == 0 {
			num = num/2
		} else {
			num = num * 3 + 1
		}
		chance++
	}
	if chance == 500 && num != 1{
		return -1
	}
	return chance
}