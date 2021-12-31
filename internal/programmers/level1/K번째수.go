import "sort"

func solution(array []int, commands [][]int) []int {
	var res []int
	for _, ijk := range commands{
		slice := append([]int{}, array[ijk[0]-1:ijk[1]]...)
		sort.Ints(slice)
		res = append(res, slice[ijk[2]-1])
	}
	return res
}