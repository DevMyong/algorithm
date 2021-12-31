package main

import "fmt"

func main() {
	var h, m uint8
	fmt.Scan(&h, &m)

	if h == 0 && m < 45 {
		fmt.Print(23, m+15)
	} else if m < 45 {
		fmt.Print(h-1, m+15)
	} else {
		fmt.Print(h, m-45)
	}
}
