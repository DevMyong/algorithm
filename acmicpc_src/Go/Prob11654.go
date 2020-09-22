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
	for _, v := range stdin.Bytes() {
		stdout.WriteString(strconv.Itoa(int(v)))
	}
	stdout.Flush()
}
