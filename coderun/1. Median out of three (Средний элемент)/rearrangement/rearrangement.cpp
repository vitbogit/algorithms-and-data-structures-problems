#include <iostream>

using namespace std;

int main(int argc, char *argv[])
{
    int a, b, c;
    cin >> a >> b >> c;
    int t;
    if (a < b){
        t = a;
        a = b;
        b = t;
    }
    // Have: a <= b
    if (a < c){
        t = a;
        a = c;
        c = t;
    }
    // Have: a <= b and a <= c
    if (b < c){
        t = b;
        b = c;
        c = t;
    }
    // Have: a <= b <= c
    cout << b;
    return 0;
}