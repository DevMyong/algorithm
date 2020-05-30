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
	stdin.Scan()
	input := strings.Fields(stdin.Text())

	newStr := make([]rune, 3)
	newNum := make([]int, 2)

	for j, word := range input {
		for i, ch := range word {
			newStr[2-i] = ch
		}
		newNum[j], _ = strconv.Atoi(string(newStr))
	}
	if newNum[0] < newNum[1] {
		stdout.WriteString(strconv.Itoa(newNum[1]))
	} else {
		stdout.WriteString(strconv.Itoa(newNum[0]))
	}
	stdout.Flush()
}
