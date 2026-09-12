#include <iostream>

using namespace std;

int main() {
    long long t;
    cin >> t;
    while (t--) {
        long long n;
        cin >> n;
        long long o = 0, z = 0;
        for (long long i = 0; i < n; ++i) {
            long long x;
            cin >> x;
            if (x == 1) {
                o++;
            } else {
                z++;
            }
        }
        if (o > z) {
            cout << "Bessie\n";
        } else if (z > o) {
            cout << "Elsie\n";
        } else {
            cout << "Bessie\n";
        }
    }
    return 0;
}