package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	stdin  = bufio.NewScanner(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())
	return n
}

func main() {
	solv2()
}
//1,2에서 속도 및 메모리는 동일함.
//문자열로 처리
func solv1() {
	mult := 1
	for i := 0; i < 3; i++ {
		num := nextInt()
		mult *= num
	}

	str := strconv.Itoa(mult)
	var num [10]int
	for i := 0; i < len(str); i++ {
		idx, _ := strconv.Atoi(str[i : i+1])
		num[idx%10]++
	}

	defer stdout.Flush()
	for i := 0; i < 10; i++ {
		stdout.WriteString(strconv.Itoa(num[i]) + "\n")
	}
}
//숫자로 처리
func solv2() {
	res := 1
	for i := 0; i < 3; i++ {
		num := nextInt()
		res *= num
	}

	var num [10]int
	for res != 0 {
		num[res%10]++
		res /= 10
	}

	defer stdout.Flush()
	for i := 0; i < 10; i++ {
		stdout.WriteString(strconv.Itoa(num[i]) + "\n")
	}
}
