#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long n;
    cin >> n;
    string s;
    cin >> s;

    long long l = 1, r = n + 2, a = -1;
    while (l <= r) {
        long long m = (l + r) / 2;
        bool ok = true;
        long long p = 0, q = 0;

        for (char c : s) {
            long long u = 1e18, v = -1e18;
            if (p == q) {
                long long w = p;
                long long x1 = w - m, y1 = w - 1;
                long long x2 = w + 1, y2 = w + m;
                if (c == '+') {
                    x1 = max(x1, 1LL); x2 = max(x2, 1LL);
                } else if (c == '-') {
                    y1 = min(y1, -1LL); y2 = min(y2, -1LL);
                } else {
                    x1 = max(x1, 0LL); y1 = min(y1, 0LL);
                    x2 = max(x2, 0LL); y2 = min(y2, 0LL);
                }
                if (x1 <= y1) { u = min(u, x1); v = max(v, y1); }
                if (x2 <= y2) { u = min(u, x2); v = max(v, y2); }
            } else {
                long long x = p - m, y = q + m;
                if (c == '+') { u = max(x, 1LL); v = y; }
                else if (c == '-') { u = x; v = min(y, -1LL); }
                else {
                    if (x <= 0 && 0 <= y) { u = 0; v = 0; }
                }
            }
            if (u > v) {
                ok = false;
                break;
            }
            p = u; q = v;
        }

        if (ok) {
            a = m;
            r = m - 1;
        } else {
            l = m + 1;
        }
    }
    cout << a << "\n";
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