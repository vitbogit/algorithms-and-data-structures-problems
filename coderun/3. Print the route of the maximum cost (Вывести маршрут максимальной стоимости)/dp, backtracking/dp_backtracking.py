N, M = map(int, input().split(" "))

dp = [[] for _ in range(N)]

for i in range(N):
    dp[i] = list(map(int, input().split(" ")))
    for j in range(M):
        from_left, from_up = -1, -1
        if i > 0:
            from_left = dp[i - 1][j]
        if j > 0:
            from_up = dp[i][j - 1]
        best = max(from_left, from_up)
        if best == -1:
            best = 0
        dp[i][j] += best
    
print(dp[N - 1][M - 1])

backtrack = []
pos_i, pos_j = N - 1, M - 1

while pos_i > 0 or pos_j > 0:
    from_left, from_up = -1, -1
    if pos_i > 0:
        from_up = dp[pos_i - 1][pos_j]
    if pos_j > 0:
        from_left = dp[pos_i][pos_j - 1]
    if from_left > from_up:
        pos_j -= 1
        backtrack += ['R']
    else:
        pos_i -= 1
        backtrack += ['D']

answer = ' '.join(backtrack[::-1])

print(answer)

