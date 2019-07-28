package main

import "fmt"

func main()  {
	fmt.Println(letterCombinations("234"))
}

//这题做得也很好！
//https://leetcode.com/problems/letter-combinations-of-a-phone-number/submissions/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
Memory Usage: 2.5 MB, less than 70.00% of Go online submissions for Letter Combinations of a Phone Number.
*/

var digitMap map[string][]string
func letterCombinations(digits string) []string {
	digitMap = make(map[string][]string)
	digitMap["2"] = []string{"a","b","c"}
	digitMap["3"] = []string{"d","e","f"}
	digitMap["4"] = []string{"g","h","i"}
	digitMap["5"] = []string{"j","k","l"}
	digitMap["6"] = []string{"m","n","o"}
	digitMap["7"] = []string{"p","q","r", "s"}
	digitMap["8"] = []string{"t","u","v"}
	digitMap["9"] = []string{"w","x","y", "z"}
	return combinationTwoString(digits)
}

func combinationTwoString(a string) []string {
	if len(a) == 1 {
		return digitMap[a]
	} else if len(a) == 2 {
		out := digitMap[a[:1]]
		in := digitMap[a[1:]]
		return combinationTwoArray(out, in)
	} else if len(a) > 2 {
		t1 := combinationTwoString(a[0:len(a) / 2])
		t2 := combinationTwoString(a[len(a) / 2:])
		return combinationTwoArray(t1, t2)
	}
	return nil
}


func combinationTwoArray(out, in []string) []string {
	t := make([]string, 0, len(out) * len(in))
	for i:=0; i < len(out); i++ {
		for j := 0; j < len(in); j++ {
			t = append(t, out[i]+in[j])
		}
	}
	return t
}