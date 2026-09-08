#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long x, y, k;
    cin >> x >> y >> k;
    if (x > y) {
        long long s = y * k + k * (k - 1) / 2;
        cout << s << "\n";
        return;
    }
    long long d = y - x;
    long long s = 0;
    long long m = d - x;
    if (m >= k) {
        m = k - 1;
    }
    for (long long i = 0; i <= m; i++) {
        s += d % (x + 1);
    }
    if (k > m + 1) {
        s += (k - (m + 1)) * d;
    }
    cout << s << "\n";
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while (t--) {
        solve();
    }
    return 0;
} 