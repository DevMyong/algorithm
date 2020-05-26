package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Scan()
	input := strings.Fields(stdin.Text())

	num := make([]int, len(input))
	for i, val := range input {
		num[i], _ = strconv.Atoi(val)
	}
	sort.Ints(num)
	stdout.WriteString(strconv.Itoa(num[1]))
	stdout.Flush()
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
