package main

import "fmt"

func main() {
	solve1()
	solve2()
}

func solve1() {
	var a, b, tmpB int
	fmt.Scan(&a, &b)
	tmpB = b
	for i := 0; i < 3; i++ {
		fmt.Println(a * (tmpB % 10))
		tmpB /= 10
	}
	fmt.Println(a * b)
}

func solve2() {
	var a, b, sum int
	sum = 0
	fmt.Scan(&a, &b)
	for i := 0; i < 3; i++ {
		fmt.Println(a * (b % 10))
		ten := 1
		for j := 0; j < i; j++ {
			ten *= 10
		}
		sum += (a * (b % 10)) * ten
		b /= 10
	}
	fmt.Println(sum)
}
