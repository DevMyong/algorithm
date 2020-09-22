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
	input := stdin.Text()

	alpha := make([]int, 26)
	for i := 0; i < 26; i++ {
		alpha[i] = -1
	}

	for i, v := range input {
		idx := int(v - 'a')
		if alpha[idx] == -1 {
			alpha[idx] = i
		}
	}
	defer stdout.Flush()
	for i := 0; i < 26; i++ {
		stdout.WriteString(strconv.Itoa(alpha[i]) + " ")
	}
}
