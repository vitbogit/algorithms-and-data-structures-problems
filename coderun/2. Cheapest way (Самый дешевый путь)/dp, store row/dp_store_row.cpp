#include <iostream>

using namespace std;

int main(int argc, char *argv[])
{
    int N, M;
    cin >> N >> M;

    int inf;
    inf = 3901;

    int* dp = new int[N];
    for (int j = 0; j < M; j++){
            int num;
            cin >> num;
            int fromLeft; 
            fromLeft = inf;
            if (j > 0) fromLeft = dp[j-1];
            int best = fromLeft;
            if (best == inf) best = 0;
            dp[j] = num + best;
    }

    for (int i = 1; i < N; i++){
        for (int j = 0; j < M; j++){
            int num;
            cin >> num;
            int fromLeft, fromAbove; 
            fromLeft = inf;
            fromAbove = inf;
            if (i > 0) fromAbove = dp[j];
            if (j > 0) fromLeft = dp[j-1];
            int best = min(fromAbove, fromLeft);
            if (best == inf) best = 0;
            dp[j] = num + best;
        }
    }

    cout << dp[M-1];

    return 0;
}