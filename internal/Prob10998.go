package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	multiply0()
	multiply1()
	multiply2()
}

func multiply0() {
	var a, b int
	fmt.Scanf("%d %d",&a,&b)
	fmt.Println(a * b)
}
func multiply1() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	split := strings.Fields(text)

	var a, b int
	a, _ = strconv.Atoi(split[0])
	b, _ = strconv.Atoi(split[1])
	fmt.Println(a * b)
}

func multiply2() {
	var (
		r = bufio.NewReader(os.Stdin)
		w = bufio.NewWriter(os.Stdout)
	)
	var a, b byte
	a, _ = r.ReadByte()
	r.ReadByte()
	b, _ = r.ReadByte()

	w.WriteString(strconv.Itoa((int(a) - 48) * (int(b) - 48)))
	w.Flush()
}
