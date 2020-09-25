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

	var str string
	for i := 1; i <= n; i++ {
		str += "*"
		wr.WriteString(str + "\n")
	}
	wr.Flush()
}
