package main

import "fmt"

func main() {

	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var addition int
			if i > 0 {
				addition = dp[i-1][j]
			}
			if j > 0 {
				addition = max(addition, dp[i][j-1])
			}

			fmt.Scan(&dp[i][j])
			dp[i][j] += addition
		}
	}

	path := []byte{} // will have extra space at the beginning

	for i, j := N-1, M-1; !(i == 0 && j == 0); {

		left := -1
		up := -1

		if i > 0 {
			up = dp[i-1][j]
		}

		if j > 0 {
			left = dp[i][j-1]
		}

		if left > up {
			path = append(path, ' ')
			path = append(path, 'R')
			j--
		} else {
			path = append(path, ' ')
			path = append(path, 'D')
			i--
		}

	}

	fmt.Println(dp[N-1][M-1])

	if len(path) > 0 {

		answer := string(reverse(path[1:]))

		fmt.Println(answer)

	}

}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
