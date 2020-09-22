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
	c, _ := strconv.Atoi(stdin.Text())

	defer stdout.Flush()
	for i := 0; i < c; i++ {
		stdin.Scan()
		input := strings.Fields(stdin.Text())
		n, _ := strconv.Atoi(input[0])

		avg := 0
		num := make([]int, n)
		for j := 0; j < n; j++ {
			num[j], _ = strconv.Atoi(input[j+1])
			avg += num[j]
		}
		avg /= n

		cnt := 0
		for k := 0; k < n; k++ {
			if num[k] > avg {
				cnt++
			}
		}
		stdout.WriteString(strconv.FormatFloat(float64(cnt)/float64(n)*100, 'f', 3, 64) + "%\n")
	}
}
