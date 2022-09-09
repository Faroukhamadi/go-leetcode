package main

func uniquePaths(m int, n int) int {
	dp := make2D[int](n, m)
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n-1][m-1]
}

func make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func main() {
	uniquePaths(3, 4)
}
