package main

import (
	"fmt"
)

func main(){
	var a, b int
	fmt.Scan(&a, &b)

	for y:=0; y<b;y++{
		for x:=0; x<a; x++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}