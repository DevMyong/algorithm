func solution(numbers []int) int {
	arr := make([]bool, 10)
	result := 0

	for _, number := range numbers {
		arr[number] = true
	}
	for i := range arr{
		if !arr[i]{
			result += i
		}
	}
	return result
}