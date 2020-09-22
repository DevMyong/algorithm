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

	defer stdout.Flush()
	for i := 0; i < n; i++ {
		var alpha [26]bool
		stdin.Scan()
		input := stdin.Text()

		if len(input) == 1 {
			cnt++
			continue
		}

		isGroupWord := true
		for i, ch := range input {
			if alpha[ch-97] && input[i-1] != byte(ch) {
				isGroupWord = false
			}
			alpha[ch-97] = true
		}
		if isGroupWord {
			cnt++
		}
	}
	stdout.WriteString(strconv.Itoa(cnt))
}
