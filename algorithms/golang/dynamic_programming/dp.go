package dynamic_programming

func uniquePaths(m int, n int) int {
	var path = 0
	if m == 1 && n == 1 {
		return 1
	}
	if n <= 0 || m <= 0 {
		return 0
	}

	if m > 1 {
		path += uniquePaths(n, m-1)
	}

	if n > 1 {
		path += uniquePaths(n-1, m)
	}
	return path
}

func uniquePaths2(m int, n int) int {
	var path [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
		path[i][0] = 1
	}

	for i := 0; i < n; i++ {
		path[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}
