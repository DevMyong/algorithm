package main

func sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += int(v)
	}
	return sum
}

func main() {

}
