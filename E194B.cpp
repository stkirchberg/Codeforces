#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long x, y, k;
    cin >> x >> y >> k;
    long long s = 0;
    for (long long i = 0; i < k; i++) {
        long long a = x + i;
        long long b = y + i;
        s += b % a;
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