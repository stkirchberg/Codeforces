#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long x, y, k;
    cin >> x >> y >> k;
    long long d = y - x;
    long long s = 0;
    long long m = min(k, max(0LL, d - x + 1));
    for (long long i = 0; i < m; i++) {
        s += (y + i) % (x + i);
    }
    if (k > m) {
        s += (k - m) * d;
    }
    cout << s << "\n";
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    long long t;
    cin >> t;
    while (t--) {
        solve();
    }
    return 0;
}