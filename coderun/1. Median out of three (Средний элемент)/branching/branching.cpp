#include <iostream>

using namespace std;

int main(int argc, char *argv[])
{
    int a, b, c;
    cin >> a >> b >> c;
    if (a < b) {
		if (b <= c) {
			cout << b;
		} else if (a <= c) {
			cout << c;
		} else {
			cout << a;
		}
	} else if (a==b) {
		cout << a;
	} else { // b < a
		if (a <= c) {
			cout << a;
		} else if (b <= c) {
			cout << c;
		} else {
			cout << b;
		}
	}
    return 0;
}