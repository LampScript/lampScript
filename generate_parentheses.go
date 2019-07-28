package main

import "fmt"



/*
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> ans = new ArrayList();
        backtrack(ans, "", 0, 0, n);
        return ans;
    }

    public void backtrack(List<String> ans, String cur, int open, int close, int max){
        if (cur.length() == max * 2) {
            ans.add(cur);
            return;
        }

        if (open < max)
            backtrack(ans, cur+"(", open+1, close, max);
        if (close < open)
            backtrack(ans, cur+")", open, close+1, max);
    }
}


//这题抄了答案还是做不对，太惭愧了

*/

func main()  {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	result := make([]string, 0, 0)
	parentheses(result, "", 0, 0, n)
	return result
}

func parentheses(result []string, s string, left, right, max int) {
	if len(s) == max * 2 {
		result = append(result, s)
		return
	}

	if left < max {
		parentheses(result, s + "(", left+1, right, max)
	}

	if right < left {
		parentheses(result, s + ")", left, right+1, max)
	}
}




