package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	stdin  = bufio.NewScanner(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
)

func main() {
	stdin.Scan()
	input := stdin.Text()

	num := len(input)

	defer stdout.Flush()
	if num == 1 {
		stdout.WriteString("1")
		return
	}
	for i := 0; i < len(input); i++ {
		if input[i] == '=' {
			if i >= 1 && (input[i-1] == 'c' || input[i-1] == 's') {
				num--
			} else if i >= 1 && input[i-1] == 'z' {
				num--
				if i >= 2 && input[i-2] == 'd' {
					num--
				}
			}

		} else if input[i] == '-' {
			num--
		} else if input[i] == 'j' && i != 0 {
			if i >= 1 && (input[i-1] == 'l' || input[i-1] == 'n') {
				num--
			}
		}
	}
	stdout.WriteString(strconv.Itoa(num))
}
