package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solv3()
}
func solv1() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Print(a+b, "\n")
	}
}
func solv2() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str, _ := r.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	n, _ := strconv.Atoi(str)

	for i := 0; i < n; i++ {
		str, _ := r.ReadString('\n')
		str = strings.TrimRight(str, "\r\n")
		strArr := strings.Split(str, " ")

		a, _ := strconv.Atoi(strArr[0])
		b, _ := strconv.Atoi(strArr[1])

		res := a + b
		w.WriteString(strconv.Itoa(res) + "\n")
	}
	w.Flush()
}
func solv3() {
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
