package main

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0{
		return 0
	}

	m:=len(obstacleGrid)
	n:=len(obstacleGrid[0])

	a := make([][]int, m)
	for i:=0;i<m;i++{
		t := make([]int, n)
		a[i] = t
		if i == 0 {
			for i2:=0;i2<len(t) ; i2++ {
				if obstacleGrid[i][i2] == 1 {
					break
				}
				t[i2]=1
			}
		}

		if obstacleGrid[i][0] == 1 {
			t[0]=0
		} else if i == 0{
			t[0]=1
		} else {
			t[0] = a[i-1][0]
		}

	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++  {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			a[i][j] = a[i-1][j]+a[i][j-1]
		}
	}
	return a[m-1][n-1]
}