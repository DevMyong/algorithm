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
	var num [42]bool
	cnt := 0

	for i := 0; i < 10; i++ {
		n := nextInt()
		idx := n % 42

		if !num[idx] {
			num[idx] = true
			cnt++
		}
	}

	stdout.WriteString(strconv.Itoa(cnt))
	stdout.Flush()
}
