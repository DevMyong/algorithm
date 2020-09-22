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

	res := 0
	for _, ch := range input {
		switch {
		case ch <= 'C':
			res += 3
		case ch <= 'F':
			res += 4
		case ch <= 'I':
			res += 5
		case ch <= 'L':
			res += 6
		case ch <= 'O':
			res += 7
		case ch <= 'S':
			res += 8
		case ch <= 'V':
			res += 9
		case ch <= 'Z':
			res += 10
		}
	}
	stdout.WriteString(strconv.Itoa(res))
	stdout.Flush()
}
