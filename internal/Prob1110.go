package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Scan()
	orgVal, _ := strconv.Atoi(stdin.Text())
	cmpVal := orgVal

	var cnt int
	for cnt = 1; ; cnt++ {
		var num [2]int

		num[0] = cmpVal / 10
		num[1] = cmpVal % 10
		cmpVal = num[1]*10 + (num[0]+num[1])%10

		if orgVal == cmpVal {
			break
		}
	}
	stdout.WriteString(strconv.Itoa(cnt))
	stdout.Flush()
}
