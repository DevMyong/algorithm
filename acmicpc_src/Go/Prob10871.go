package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	solv1()
}

func solv1() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	sc.Scan()
	nx := strings.Fields(sc.Text())
	n, _ := strconv.Atoi(nx[0])
	x, _ := strconv.Atoi(nx[1])

	sc.Scan()
	strArr := strings.Fields(sc.Text())
	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(strArr[i])
		if val < x {
			wr.WriteString(strArr[i] + " ")
		}
	}
	wr.Flush()
}
