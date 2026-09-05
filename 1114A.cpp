#include <iostream>
using namespace std;

int main() {
    int t;
    cin >> t;

    for (int i = 0; i < t; i++) {
        long long a, b, c;
        cin >> a >> b >> c;

        int rounds = 0;

        while (a != b && a != c && b != c) {

            if (a > b && a > c) {
                a--;
            } else if (b > a && b > c) {
                b--;
            } else {
                c--;
            }

            if (a < b && a < c) {
                a++;
            } else if (b < a && b < c) {
                b++;
            } else {
                c++;
            }

            rounds++;
        }

        cout << rounds << endl;
    }

    return 0;
}