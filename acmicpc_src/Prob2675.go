package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	stdin  = bufio.NewScanner(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
)

func main() {
	stdin.Scan()
	tc, _ := strconv.Atoi(stdin.Text())

	defer stdout.Flush()
	for i := 0; i < tc; i++ {
		stdin.Scan()
		input := strings.Fields(stdin.Text())

		r, _ := strconv.Atoi(input[0])
		var newStr string

		for _, ch := range input[1] {
			newStr += strings.Repeat(string(ch), r)
		}
		stdout.WriteString(newStr + "\n")
	}
}
