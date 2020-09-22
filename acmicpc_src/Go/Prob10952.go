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

	for sc.Scan() {
		if sc.Text() == "0 0" {
			break
		}
		strArr := strings.Fields(sc.Text())
		a, _ := strconv.Atoi(strArr[0])
		b, _ := strconv.Atoi(strArr[1])
		wr.WriteString(strconv.Itoa(a+b) + "\n")
	}
	wr.Flush()
}
