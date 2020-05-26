package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	for i := 0; i < 2*n-1; i++ {
		if i < n {
			stdout.WriteString(strings.Repeat("*", i+1) + "\n")
		} else {
			stdout.WriteString(strings.Repeat("*", 2*n-i-1) + "\n")
		}
	}
	stdout.Flush()
}
