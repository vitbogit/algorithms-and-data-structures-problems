inf = 300000

N = int(input())

dp = [[inf] * (N+2) for _ in range(N+1)]

dp[0][0] = 0

cost = [0] * N

for i in range(N):
    cost[i] = int(input())

for i in range(1, N+1):
    for j in range(N+1):
        if cost[i-1] <= 100:
            dp[i][j] = min(dp[i-1][j] + cost[i-1], dp[i-1][j+1])
        else:
            if j > 0:
                dp[i][j] = min(dp[i-1][j-1] + cost[i-1], dp[i-1][j+1])
            else:
                dp[i][j] = dp[i-1][j+1]

answ = dp[N][0]
k1, k2 = 0, 0

for i in range(N+1):
    if dp[N][i] <= answ:
        answ = dp[N][i]
        k1 = i


#print("dp:", dp, "/dp")


path = []

i, j = N, k1
while i > 0:
    if cost[i-1] <= 100:
        if dp[i][j] == dp[i-1][j+1]:
            path += [i]
            j+=1
    else:
        if j > 0:
            if dp[i][j] == dp[i-1][j+1]:
                path += [i]
                j+=1
            else:
                j-=1
        else:
            path += [i]
            j+=1
    i-=1

k2 = len(path)

path.reverse()

print(answ)
print(k1, k2)
for day in path:
    print(day)