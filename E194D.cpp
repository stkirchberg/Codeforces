#include <bits/stdc++.h>
using namespace std;

bool f(long long k, string s) {
    long long l = 0, r = 0;
    for (char c : s) {
        long long u = 1e18, v = -1e18;
        if (l == r) {
            long long w = l;
            long long a1 = w - k, b1 = w - 1;
            long long a2 = w + 1, b2 = w + k;
            if (c == '+') {
                a1 = max(a1, 1LL); a2 = max(a2, 1LL);
            } else if (c == '-') {
                b1 = min(b1, -1LL); b2 = min(b2, -1LL);
            } else {
                a1 = max(a1, 0LL); b1 = min(b1, 0LL);
                a2 = max(a2, 0LL); b2 = min(b2, 0LL);
            }
            if (a1 <= b1) { u = min(u, a1); v = max(v, b1); }
            if (a2 <= b2) { u = min(u, a2); v = max(v, b2); }
        } else {
            long long a = l - k, b = r + k;
            if (c == '+') { u = max(a, 1LL); v = b; }
            else if (c == '-') { u = a; v = min(b, -1LL); }
            else {
                if (a <= 0 && 0 <= b) { u = 0; v = 0; }
            }
        }
        if (u > v) return false;
        l = u; r = v;
    }
    return true;
}

void solve() {
    long long n;
    cin >> n;
    string s;
    cin >> s;
    long long l = 1, r = n + 2, a = -1;
    while (l <= r) {
        long long m = (l + r) / 2;
        if (f(m, s)) {
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