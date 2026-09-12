#include <bits/stdc++.h>
using namespace std;

long long g = 1000000007;

void h() {
    long long n;
    if (!(cin >> n)) return;
    vector<long long> a(n + 1);
    for (long long i = 1; i <= n; ++i) cin >> a[i];

    vector<long long> f(n + 2, 0);
    vector<long long> m(n + 1, -1);

    for (long long i = 1; i <= n; ++i) {
        long long l = a[i] * i;
        long long r = (a[i] + 1) * i - 1;
        if (l < n) {
            f[l]++;
            f[min(r, n - 1) + 1]--;
        }
        for (long long v = 0; v < a[i]; ++v) {
            long long x = v * i;
            long long y = (v + 1) * i - 1;
            if (x < n) {
                long long z = min(y, n - 1);
                m[z + 1] = max(m[z + 1], x);
            }
        }
    }

    long long c = 0;
    vector<long long> b(n, 0);
    for (long long i = 0; i < n; ++i) {
        c += f[i];
        if (c > 0) b[i] = 1;
    }

    for (long long i = 1; i <= n; ++i) {
        m[i] = max(m[i], m[i - 1]);
    }

    vector<long long> d(n + 2, 0);
    vector<long long> s(n + 2, 0);

    d[0] = 1;
    s[0] = 1;

    for (long long i = 0; i <= n; ++i) {
        long long u = max(0LL, m[i] + 1);
        long long w = i;
        if (u <= w) {
            long long p = (s[w] - (u > 0 ? s[u - 1] : 0) + g) % g;
            if (i == n || !b[i]) {
                d[i + 1] = p;
            }
        }
        s[i + 1] = (s[i] + d[i + 1]) % g;
    }

    cout << d[n + 1] << "\n";
}

int main() {
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    long long t;
    if (cin >> t) {
        while (t--) {
            h();
        }
    }
    return 0;
}