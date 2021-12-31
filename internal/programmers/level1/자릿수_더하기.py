import "fmt"

func solution(n int) int {
	answer := 0

	for n > 0 {
		answer += n%10
		n /= 10
	}
	// [실행] 버튼을 누르면 출력 값을 볼 수 있습니다.
	fmt.Printf("Hello Go")

	return answer
}