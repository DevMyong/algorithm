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

	defer stdout.Flush()
	for i := 0; i < n; i++ {
		stdin.Scan()
		input := stdin.Text()

		score, cnt := 0, 0
		for _, ch := range input {
			if ch == 'O' {
				cnt++
				score += cnt
			} else {
				cnt = 0
			}
		}
		stdout.WriteString(strconv.Itoa(score) + "\n")
	}
}
