#include <bits/stdc++.h>
using namespace std;

bool f(int k, string& s, int n) {
    long long l = 0, r = 0;
    for (int i = 0; i < n; i++) {
        char c = s[i];
        if (c == '0') {
            bool w = (max(l, -(long long)k) <= min(r, -1LL)) || (max(l, 1LL) <= min(r, (long long)k));
            if (!w) return 0;
            l = 0; r = 0;
        } else if (c == '+') {
            long long u = r + k;
            if (u < 1) return 0;
            long long v = (r >= 2) ? max(1LL, l - k) : max(1LL, l + 1);
            l = v; r = u;
        } else {
            long long v = l - k;
            if (v > -1) return 0;
            long long u = (l <= -2) ? min(-1LL, r + k) : min(-1LL, r - 1);
            l = v; r = u;
        }
    }
    return 1;
}

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
    int t;
    if (!(cin >> t)) return 0;
    while (t--) {
        int n;
        string s;
        cin >> n >> s;
        int a = 1, b = n, z = -1;
        while (a <= b) {
            int m = (a + b) / 2;
            if (f(m, s, n)) {
                z = m;
                b = m - 1;
            } else {
                a = m + 1;
            }
        }
        cout << z << "\n";
    }
    return 0;
}