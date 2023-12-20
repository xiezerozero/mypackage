package demo1

func MaxValue(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var m = len(grid)    //横向
	var n = len(grid[0]) //纵向

	var dp = make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1, n+1))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[m][n]
}

func MaxValue2(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var m = len(grid)    //横向
	var n = len(grid[0]) //纵向
	var dp = make([]int, n+1, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[j+1] = max(dp[j], dp[j+1]) + grid[i][j]
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
