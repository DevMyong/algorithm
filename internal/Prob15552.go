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
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		strArr := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(strArr[0])
		b, _ := strconv.Atoi(strArr[1])

		wr.WriteString(strconv.Itoa(a+b) + "\n")
	}
	wr.Flush()
}
