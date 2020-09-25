package pkg

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