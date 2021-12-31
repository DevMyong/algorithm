package main

import (
	"bufio"
	"os"
	"strconv"
)

// var (
// 	stdin  = bufio.NewScanner(os.Stdin)
// 	stdout = bufio.NewWriter(os.Stdout)
// )

// func nextInt() int {
// 	stdin.Scan()
// 	n, _ := strconv.Atoi(stdin.Text())
// 	return n
// }

func main() {
	solv2()
}

// 함수 나누어서 진행 (메모리 8 증가)
func solv1() {
	// max, idx := 0, 0

	// for i := 0; i < 9; i++ {
	// 	num := nextInt()

	// 	if num > max {
	// 		max, idx = num, i
	// 	}
	// }
	// stdout.WriteString(strconv.Itoa(max) + "\n" + strconv.Itoa(idx+1))
	// stdout.Flush()
}

func solv2() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	max, idx := 0, 0

	defer stdout.Flush()
	for i := 0; i < 9; i++ {
		stdin.Scan()
		num, _ := strconv.Atoi(stdin.Text())
		if num > max {
			max, idx = num, i+1
		}
	}
	stdout.WriteString(strconv.Itoa(max) + "\n" + strconv.Itoa(idx))
}
