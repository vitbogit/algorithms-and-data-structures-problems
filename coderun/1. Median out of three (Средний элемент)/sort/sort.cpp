#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(int argc, char *argv[])
{
    int a, b, c;
    cin >> a >> b >> c;
    int arr[] = {a, b, c};
    sort(arr, arr + 3);
    cout << arr[1];
    return 0;
}