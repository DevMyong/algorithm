package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

var (
	stdin  = bufio.NewScanner(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
)

func init() {
	const BUFSIZE = 1024 * 1000
	stdin.Buffer(make([]byte, BUFSIZE), BUFSIZE)
}

func main() {
	solv2()
}
func solv1() {
	stdin.Scan()
	input := strings.Fields(strings.TrimSpace(stdin.Text()))
	stdout.WriteString(strconv.Itoa(len(input)))
	stdout.Flush()
}
func solv2() {
	stdin.Scan()
	input := bytes.TrimSpace(stdin.Bytes())

	if len(input) == 0 {
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
