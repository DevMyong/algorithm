package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	var sum int
	for i := 0; i < 5; i++ {
		stdin.Scan()
		input := stdin.Text()
		num, _ := strconv.Atoi(input)
		if num < 40 {
			sum += 40
		} else {
			sum += num
		}

	}
	stdout.WriteString(strconv.Itoa(sum / 5))
	stdout.Flush()
}
