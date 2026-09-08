#include <bits/stdc++.h>
using namespace std;

bool c(long long k, const string& s) {
    long long l = 0, r = 0;
    for (char x : s) {
        long long u, v;
        if (l == r) {
            long long w = l;
            long long a = w - k, b = w - 1;
            long long d = w + 1, e = w + k;
            if (x == '+') {
                a = max(a, 1LL);
                d = max(d, 1LL);
            } else if (x == '-') {
                b = min(b, -1LL);
                e = min(e, -1LL);
            } else {
                a = max(a, 0LL); b = min(b, 0LL);
                d = max(d, 0LL); e = min(e, 0LL);
            }
            bool f = (a <= b);
            bool g = (d <= e);
            if (f && g) {
                u = min(a, d);
                v = max(b, e);
            } else if (f) {
                u = a; v = b;
            } else if (g) {
                u = d; v = e;
            } else {
                return false;
            }
        } else {
            long long a = l - k, b = r + k;
            if (x == '+') {
                u = max(a, 1LL);
                v = b;
            } else if (x == '-') {
                u = a;
                v = min(b, -1LL);
            } else {
                if (a <= 0 && 0 <= b) {
                    u = 0; v = 0;
                } else {
                    return false;
                }
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
    
    long long a = 1, b = n + 2, p = -1;
    while (a <= b) {
        long long m = a + (b - a) / 2;
        if (c(m, s)) {
            p = m;
            b = m - 1;
        } else {
            a = m + 1;
        }
    }
    cout << p << "\n";
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