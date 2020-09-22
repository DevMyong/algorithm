package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sub0()
	sub1()
	sub2()
}

func sub0() {
	var a, b int8
	fmt.Scan(&a, &b)
	fmt.Println(a - b)
}
func sub1() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	split := strings.Fields(text)

	var a, b int
	a, _ = strconv.Atoi(split[0])
	b, _ = strconv.Atoi(split[1])
	fmt.Println(a - b)
}

func sub2() {
	var (
		r = bufio.NewReader(os.Stdin)
		w = bufio.NewWriter(os.Stdout)
	)
	var a, b byte
	a, _ = r.ReadByte()
	r.ReadByte()
	b, _ = r.ReadByte()

	w.WriteString(strconv.Itoa(int(a) - int(b)))
	w.Flush()
}
