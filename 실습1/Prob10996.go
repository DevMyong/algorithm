package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	solv2()
}

func solv1() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	defer stdout.Flush()
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				if j%2 == 0 {
					stdout.WriteString("*")
				} else {
					stdout.WriteString(" ")
				}
			}
			stdout.WriteString("\n")
		} else if n != 1 {
			for j := 0; j < n; j++ {
				if j%2 == 0 {
					stdout.WriteString(" ")
				} else {
					stdout.WriteString("*")
				}
			}
			stdout.WriteString("\n")
		}
	}
}
func solv2() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	defer stdout.Flush()

	if n == 1 {
		stdout.WriteByte('*')
		return
	}
	for i := 0; i < 2*n; i++ {
		for j := 0; j < n; j++ {
			if (i+j)%2 == 0 {
				stdout.WriteString("*")
			} else {
				stdout.WriteString(" ")
			}
		}
		stdout.WriteString("\n")
	}
}
