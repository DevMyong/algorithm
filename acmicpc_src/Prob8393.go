package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {

}

func solv1(){
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	wr.WriteString(strconv.Itoa(n*(n+1)/2) + "\n")
	wr.Flush()
}
func solv2(){
	var n int
	fmt.Scan(&n)
	fmt.Print(n*(n+1)/2)
}
