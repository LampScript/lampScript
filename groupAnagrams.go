package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "a名称"
	fmt.Println(len(str))
	fmt.Println(len([]byte(str)))
	fmt.Println(len([]rune(str)))
}

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string, len(strs))

	array := make([][]string, 0)

	for _, v := range strs {

		t := make([]int, 0, len(v))
		for _,r := range v{
			t = append(t, int(r))
		}
		sort.Ints(t)
		z := ""
		for _,r := range t{
			z += string(r)
		}

		if str, ok := m[z]; !ok {
			str = make([]string, 0)
			str = append(str, v)
			m[z] = str
		} else {
			str = append(str, v)
			m[z] = str
		}
	}

	for _,v := range m {
		array = append(array, v)
	}
	return array
}