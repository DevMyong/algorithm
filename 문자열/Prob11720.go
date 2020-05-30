package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	stdin  = bufio.NewScanner(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
)

func main() {
	solv2()
}
func solv1() {
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	stdin.Scan()
	input := stdin.Text()
	sum := 0
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(input[i : i+1])
		sum += num
	}

	stdout.WriteString(strconv.Itoa(sum))
	stdout.Flush()
}
func solv2() {
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	stdin.Scan()
	input := strings.Split(stdin.Text(), "")

	sum := 0
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(input[i])
		sum += num
	}
	stdout.WriteString(strconv.Itoa(sum))
	stdout.Flush()
}
