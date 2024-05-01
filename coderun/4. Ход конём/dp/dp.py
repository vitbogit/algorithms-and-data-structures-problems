N, M = map(int, input().split())

dp = [[0] * (M) for _ in range(N)]

dp[0][0] = 1

for i in range(N):
    for j in range(M):
        x1, y1 = i+1, j+2
        if x1 < N and y1 < M:
            dp[x1][y1] += dp[i][j]
        x1, y1 = i+2, j+1
        if x1 < N and y1 < M:
            dp[x1][y1] += dp[i][j]

print(dp[N-1][M-1])