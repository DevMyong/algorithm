import "sort"

type fuckinString struct{
	strings string
	chIndex string
}

func solution(strings []string, n int) []string {
	sort.Strings(strings)

	fs := make([]fuckinString, 0)


	for _, str := range strings {
		fs = append(fs, fuckinString {str, string(str[n]) })
	}

	sort.SliceStable(fs, func(i, j int) bool { return fs[i].chIndex < fs[j].chIndex })
	println(fs)
	
	for i, str := range fs {
		strings[i] = str.strings
	}
	
	return strings
}