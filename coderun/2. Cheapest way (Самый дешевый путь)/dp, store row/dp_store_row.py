N, M = map(int, input().split(" "))

dp = [0] * (M)

#print("echo dp: ", dp)

inf = 3901

# first row
row = list(map(int, input().split(" ")))
#print("echo row: ", row)
for j in range(M):
    from_left = inf
    if j > 0:
        from_left = dp[j - 1]
    best = from_left
    if best == inf:
        best = 0
    #print("j:", j)
    dp[j] = best + row[j]

for i in range(N-1):
    row = list(map(int, input().split(" ")))
    for j in range(M):
        from_left, from_up = inf, inf
        from_up = dp[j]
        if j > 0:
            from_left = dp[j - 1]
        best = min(from_left, from_up)
        if best == inf:
            best = 0
        dp[j] = best + row[j]

print(dp[M - 1])