#include <iostream>
#include <string>
#include <algorithm> 

using namespace std;

int main(int argc, char *argv[]){
    int N, M;
    cin >> N >> M;
    
    int **dp;
    dp = new int*[N];
    for (int i = 0; i < N; i++){
        dp[i] = new int[M];
        for (int j = 0; j < M; j++){
            cin >> dp[i][j];
            int fromLeft, fromAbove;
            fromLeft = -1;
            fromAbove = -1;
            if (i > 0){
                fromAbove = dp[i-1][j];
            }
            if (j > 0){
                fromLeft = dp[i][j-1];
            }
            int best;
            best = max(fromAbove, fromLeft);
            if (best == -1){
                best = 0;
            }
            dp[i][j] += best;
        }
    }
    
    cout << dp[N-1][M-1] << endl;

    string answ = "";

    int pos_x, pos_y;
    pos_x = N-1;
    pos_y = M-1;

    while (pos_x > 0 || pos_y > 0){
        int fromLeft, fromAbove;
        fromLeft = -1;
        fromAbove = -1;
        if (pos_x > 0){
            fromAbove = dp[pos_x-1][pos_y];
        }
        if (pos_y > 0){
            fromLeft = dp[pos_x][pos_y-1];
        }
        if (fromAbove > fromLeft){
            answ += "D ";
            pos_x--;
        } else {
            answ += "R ";
            pos_y--;
        }
    }

    if(answ.length() > 0){
        answ = answ.substr(0, answ.length()-1);
    }
    reverse(answ.begin(), answ.end());

    cout << answ;

    return 0;
}