package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	cost := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&cost[i])
	}

	inf := 300000

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, N+2)
		for j := 0; j < N+2; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := 1; i < N+1; i++ {
		for j := 0; j < N+1; j++ {
			if cost[i-1] <= 100 {
				//fmt.Println("move a")
				dp[i][j] = min(dp[i-1][j]+cost[i-1], dp[i-1][j+1])
			} else {
				if j > 0 {
					//fmt.Println("move b")
					dp[i][j] = min(dp[i-1][j-1]+cost[i-1], dp[i-1][j+1])
				} else {
					//fmt.Println("move c")
					dp[i][j] = dp[i-1][j+1]
				}
			}
		}
	}

	k1, k2 := 0, 0
	answ := dp[N][0]

	for i := 0; i < N+1; i++ {
		if dp[N][i] <= answ {
			answ = dp[N][i]
			k1 = i
		}
	}

	path := []int{}

	for i, j := N, k1; i > 0; i-- {
		if cost[i-1] <= 100 {
			if dp[i][j] == dp[i-1][j+1] {
				path = append(path, i)
				j++
			}
		} else {
			if j > 0 {
				if dp[i][j] == dp[i-1][j-1]+cost[i-1] {
					j--
				} else {
					path = append(path, i)
					j++
				}
			} else {
				path = append(path, i)
				j++
			}
		}
	}

	reverse(path)

	k2 = len(path)

	fmt.Println(answ)
	fmt.Printf("%d %d\n", k1, k2)
	for _, day := range path {
		fmt.Println(day)
	}

	//fmt.Println(cost)
	//fmt.Println(dp)

}

func reverse(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
