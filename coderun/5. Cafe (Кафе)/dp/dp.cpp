#include <iostream>
#include <vector>
#include <algorithm>

const int INF = 300000;

int main() {
    int N;
    std::cin >> N;

    std::vector<std::vector<int>> dp(N + 1, std::vector<int>(N + 2, INF));
    dp[0][0] = 0;

    std::vector<int> cost(N);
    for (int i = 0; i < N; ++i) {
        std::cin >> cost[i];
    }

    for (int i = 1; i <= N; ++i) {
        for (int j = 0; j <= N; ++j) {
            if (cost[i - 1] <= 100) {
                dp[i][j] = std::min(dp[i - 1][j] + cost[i - 1], dp[i - 1][j + 1]);
            } else {
                if (j > 0) {
                    dp[i][j] = std::min(dp[i - 1][j - 1] + cost[i - 1], dp[i - 1][j + 1]);
                } else {
                    dp[i][j] = dp[i - 1][j + 1];
                }
            }
        }
    }

    int answ = dp[N][0];
    int k1 = 0, k2 = 0;

    for (int i = 0; i <= N; ++i) {
        if (dp[N][i] <= answ) {
            answ = dp[N][i];
            k1 = i;
        }
    }

    std::vector<int> path;

    int i = N, j = k1;
    while (i > 0) {
        if (cost[i - 1] <= 100) {
            if (dp[i][j] == dp[i - 1][j + 1]) {
                path.push_back(i);
                j++;
            }
        } else {
            if (j > 0) {
                if (dp[i][j] == dp[i - 1][j + 1]) {
                    path.push_back(i);
                    j++;
                } else {
                    j--;
                }
            } else {
                path.push_back(i);
                j++;
            }
        }
        i--;
    }

    k2 = path.size();
    std::reverse(path.begin(), path.end());

    std::cout << answ << std::endl;
    std::cout << k1 << " " << k2 << std::endl;
    for (const int& day : path) {
        std::cout << day << std::endl;
    }

    return 0;
}
