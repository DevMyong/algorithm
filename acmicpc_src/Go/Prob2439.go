package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 1; i <= n; i++ {
		wr.WriteString(strings.Repeat(" ", n-i) + strings.Repeat("*", i) + "\n")
	}
	wr.Flush()
}
