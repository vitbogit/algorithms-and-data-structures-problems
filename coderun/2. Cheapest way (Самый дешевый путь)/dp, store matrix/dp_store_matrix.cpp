#include <iostream>

using namespace std;

int main(int argc, char *argv[])
{
    int N, M;
    cin >> N >> M;

    int inf;
    inf = 3901;

    int** dp = new int*[N];
    for (int i = 0; i < N; i++){
        dp[i] = new int[M];
        for (int j = 0; j < M; j++){
            int num;
            cin >> num;
            int fromLeft, fromAbove; 
            fromLeft = inf;
            fromAbove = inf;
            if (i > 0) fromAbove = dp[i-1][j];
            if (j > 0) fromLeft = dp[i][j-1];
            int best = min(fromAbove, fromLeft);
            if (best == inf) best = 0;
            dp[i][j] = num + best;
        }
    }

    cout << dp[N-1][M-1];

    return 0;
}