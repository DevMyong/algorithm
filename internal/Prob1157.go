package main

import (
	"bufio"
	"os"
)

var (
	stdin  = bufio.NewScanner(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
)

func init() {
	const BUFSIZE = 1024 * 1000
	buf := make([]byte, BUFSIZE)
	stdin.Buffer(buf, BUFSIZE)
}

func main() {

}

func solv1() {
	stdin.Scan()
	input := stdin.Text()

	alpha := make([]int, 26)
	for _, c := range input {
		if c <= 'Z' {
			c = c + 32
		}
		idx := int(c - 'a')
		alpha[idx]++
	}

	var ch byte
	max := 0
	isEqual := false

	for i, c := range alpha {
		if c > max {
			max = alpha[i]
			ch = byte(i)
			isEqual = false
		} else if alpha[i] == max && max != 0 {
			isEqual = true
		}
	}

	if isEqual {
		stdout.WriteString("?")
	} else {
		stdout.WriteString(string('A' + ch))
	}
	stdout.Flush()
}
