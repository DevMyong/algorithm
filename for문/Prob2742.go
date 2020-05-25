package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := n; i >= 1; i-- {
		wr.WriteString(strconv.Itoa(i) + "\n")
	}
	wr.Flush()
}
