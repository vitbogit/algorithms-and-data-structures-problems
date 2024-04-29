package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	row := make([]int, M)

	// First row
	fmt.Scan(&row[0])
	for i := 1; i < M; i++ {
		fmt.Scan(&row[i])
		row[i] += row[i-1]
	}

	// Remaining rows
	var num int // tmp variable
	for j := 1; j < N; j++ {

		// First element of row
		fmt.Scan(&num)
		row[0] += num

		// Remaining elements of row
		for i := 1; i < M; i++ {
			fmt.Scan(&num)
			row[i] = num + min(row[i-1], row[i])
		}
	}

	fmt.Println(row[M-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
