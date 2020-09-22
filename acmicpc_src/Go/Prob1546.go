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

func init() {
	stdin.Split(bufio.ScanWords)
}
func nextInt() int {
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())
	return n
}

func main() {
	n := nextInt()

	num := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		num[i] = nextInt()
		if num[i] > max {
			max = num[i]
		}
	}
	var sum float64
	for i := 0; i < n; i++ {
		sum += float64(num[i]) / float64(max) * 100
	}

	stdout.WriteString(strconv.FormatFloat(sum/float64(n), 'f', -1, 64))
	stdout.Flush()
}
