package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	arr := make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = make([]int, M)
	}

	arr[0][0] = 1

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			x1, y1 := i+1, j+2
			if x1 < N && y1 < M {
				arr[x1][y1] += arr[i][j]
			}
			x1, y1 = i+2, j+1
			if x1 < N && y1 < M {
				arr[x1][y1] += arr[i][j]
			}
		}
	}

	fmt.Println(arr[N-1][M-1])
}
