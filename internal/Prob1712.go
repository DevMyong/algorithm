package main

import (
	"bufio"
	"os"
	"strconv"
	"bytes"
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
	// static A$, dynamic B$, Price C$
	// A+B < C (it's break-even point)
	stdin.Scan()
	input := bytes.TrimSpace(stdin.Bytes())

	if len(input) == 0{
		stdout.WriteString("0")
		stdout.Flush()
		return
	}
	cnt := 1
	for _, ch := range input {
		if ch == ' ' {
			cnt++
		}
	}
	stdout.WriteString(strconv.Itoa(cnt))
	stdout.Flush()
}