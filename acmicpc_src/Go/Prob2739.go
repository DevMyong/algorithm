package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 10; i++ {
		fmt.Print(n, " * ", i, " = ", n*i, "\n")
	}
}
