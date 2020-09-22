package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

/* Scanner buffer의 default size는 65536 이다.
 * Buffer나 Split 메소드를 사용하여 크기, 단위 등을 조절한다.
 */
func main() {
	solv3()
}

//with sort
func solv1() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Split(bufio.ScanWords)
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	num := make([]int, n)
	for i := 0; i < n; i++ {
		stdin.Scan()
		num[i], _ = strconv.Atoi(stdin.Text())
	}
	sort.Ints(num)
	stdout.WriteString(strconv.Itoa(num[0]) + " " + strconv.Itoa(num[n-1]))
	stdout.Flush()
}

//
func solv2() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	stdin.Split(bufio.ScanWords)
	stdin.Scan()
	n, _ := strconv.Atoi(stdin.Text())

	min := 1000001
	max := -1000001
	for i := 0; i < n; i++ {
		stdin.Scan()
		num, _ := strconv.Atoi(stdin.Text())
		if num <= min {
			min = num
		}
		if num >= max {
			max = num
		}
	}
	stdout.WriteString(strconv.Itoa(min) + " " + strconv.Itoa(max))
	stdout.Flush()
}
