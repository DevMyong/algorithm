package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	solv1()
	solv2()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func solv1() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	var cost [5]int

	for i := 0; i < 5; i++ {
		stdin.Scan()
		cost[i], _ = strconv.Atoi(stdin.Text())
	}

	minCost := cost[0]
	for i := 1; i < 3; i++ {
		if minCost > cost[i] {
			minCost = cost[i]
		}
	}
	if cost[3] < cost[4] {
		minCost += cost[3] - 50
	} else {
		minCost += cost[4] - 50
	}

	stdout.WriteString(strconv.Itoa(minCost))
	stdout.Flush()
}

func solv2(){
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	var cost [5]int
	for i := 0; i < 5; i++ {
		stdin.Scan()
		cost[i], _ = strconv.Atoi(stdin.Text())
	}

	stdout.WriteString(strconv.Itoa(min(cost[0], min(cost[1], cost[2])) + min(cost[3], cost[4]) - 50))
	stdout.Flush()
}
