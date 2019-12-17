package main

import "fmt"

func main() {

	fmt.Println(uniquePaths(51,9))
}



func uniquePaths(m int, n int) int {
	a := make([][]int, m)
	for i:=0;i<m;i++{
		t := make([]int, n)
		if i == 0 {
			for i2:=0;i2<len(t) ; i2++ {
				t[i2]=1
			}
		}
		t[0]=1
		a[i] = t
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++  {
			a[i][j] = a[i-1][j]+a[i][j-1]
		}
	}
	return a[m-1][n-1]
}