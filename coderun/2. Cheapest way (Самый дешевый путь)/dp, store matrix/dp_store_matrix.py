N, M = map(int, input().split(" "))

dp = [[0] * (M) for _ in range(N)]

inf = 3901

for i in range(N):
    row = list(map(int, input().split(" ")))
    for j in range(M):
        from_left, from_up = inf, inf
        if i > 0:
            from_left = dp[i - 1][j]
        if j > 0:
            from_up = dp[i][j - 1]
        best = min(from_left, from_up)
        if best == inf:
            best = 0
        dp[i][j] = best + row[j]

print(dp[N - 1][M - 1])