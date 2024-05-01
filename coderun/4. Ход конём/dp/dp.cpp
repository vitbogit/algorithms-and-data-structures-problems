#include <iostream>
#include <string>
#include <algorithm> 

using namespace std;

int main(int argc, char *argv[]){
    int N, M;
    cin >> N >> M;

    int** dp = new int*[N];
    for(int i = 0; i < N; i++){
        dp[i] = new int[M];
        for(int j = 0; j < M; j++){
            dp[i][j] = 0;
        }
    }

    dp[0][0] = 1;

    for(int i = 0; i < N; i++){
        for(int j = 0; j < M; j++){
            int x1, y1;
            x1 = i+1;
            y1 = j+2;
            if(x1 < N && y1 < M){
                dp[x1][y1] += dp[i][j];
            }
            x1 = i+2;
            y1 = j+1;
            if(x1 < N && y1 < M){
                dp[x1][y1] += dp[i][j];
            }
        }
    }

    cout << dp[N-1][M-1];

}