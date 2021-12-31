func solution(a int, b int) string {
	month := []int{1,-1,1,0,1,0,1,1,0,1,0,1}
	days := b
	for i:=0; i<a-1; i++{
		days += month[i] + 30
	}
	weeks := []string{1:"FRI",2:"SAT",3:"SUN",4:"MON",5:"TUE",6:"WED",0:"THU"}

	return weeks[days%7]
}