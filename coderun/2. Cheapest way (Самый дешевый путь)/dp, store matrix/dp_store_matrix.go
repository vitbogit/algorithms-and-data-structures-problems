package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	inf := 39001 // 19 numbers on way, each 100

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var x int
			fmt.Scan(&x)
			fromLeft, fromAbove := inf, inf
			if i > 0 {
				fromAbove = dp[i-1][j]
			}
			if j > 0 {
				fromLeft = dp[i][j-1]
			}
			best := min(fromLeft, fromAbove)
			if best == inf {
				best = 0
			}
			dp[i][j] = x + best
		}
	}

	fmt.Print(dp[N-1][M-1])
}
