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

func main() {
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	cnt := 0
	for i := 1; i <= n; i++ {
		if solv2(i) {
			cnt++
		}
	}
	stdout.WriteString(strconv.Itoa(cnt))
	stdout.Flush()
}

func solv1(n int) {
	cnt := 0

	for k := 1; k <= n; k++ {
		var num [4]int
		tmp := k
		for i := 0; tmp > 0; i++ {
			num[i] = tmp % 10
			tmp /= 10
		}
		if k < 100 {
			cnt++
		} else if k < 1000 && (num[0]-num[1]) == (num[1]-num[2]) {
			cnt++
		}
	}
	stdout.WriteString(strconv.Itoa(cnt))
}

func solv2(n int) bool {
	if n < 100 {
		return true
	} else if n < 1000 {
		var num [3]int
		for i := 0; i < 3; i++ {
			num[i] = n % 10
			n /= 10
		}
		if num[0]-num[1] == num[1]-num[2] {
			return true
		}
	}
	return false
}
