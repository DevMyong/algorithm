package main

import (
	"bufio"
	"os"
	"strconv"
)

var stdout = bufio.NewWriter(os.Stdout)

func main() {
	var selfNum [10000]bool
	for i := 1; i < 10000; i++ {
		idx := selfNumber(i)
		if idx < 10000 && !selfNum[idx] {
			selfNum[idx] = true
		}
	}
	for i := 1; i < 10000; i++ {
		if !selfNum[i] {
			stdout.WriteString(strconv.Itoa(i) + "\n")
		}
	}
	stdout.Flush()
}

func selfNumber(a int) int {
	res := a
	for a > 0 {
		res += a % 10
		a = a / 10
	}
	return res
}
