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

	defer wr.Flush()
	for sc.Scan() {
		strArr := strings.Fields(sc.Text())
		a, _ := strconv.Atoi(strArr[0])
		b, _ := strconv.Atoi(strArr[1])
		wr.WriteString(strconv.Itoa(a+b) + "\n")
	}
	// strArr -> input
	// sc ->stdin, wr -> stdout
}
