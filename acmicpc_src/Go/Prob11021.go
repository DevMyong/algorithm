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
		sc.Scan()
		strArr := strings.Fields(sc.Text())
		a, _ := strconv.Atoi(strArr[0])
		b, _ := strconv.Atoi(strArr[1])

		str := "Case #" + strconv.Itoa(i) + ": " + strconv.Itoa(a+b)
		wr.WriteString(str + "\n")
	}
	wr.Flush()
}
