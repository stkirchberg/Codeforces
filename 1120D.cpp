#include <bits/stdc++.h>
using namespace std;

long long x[1000005];
long long y[1000005];
long long z[200005]; 

void g(long long o, long long l, long long r, long long i, long long v) {
    if (l == r) {
        x[o] = v;
        y[o] = v;
        return;
    }
    long long d = (l + r) / 2;
    if (i <= d) g(2 * o, l, d, i, v);
    else g(2 * o + 1, d + 1, r, i, v);
    x[o] = max(x[2 * o], x[2 * o + 1]);
    y[o] = y[2 * o] + y[2 * o + 1];
}

long long q(long long o, long long l, long long r, long long h, long long j) {
    if (h > j) return 0;
    if (l >= h && r <= j) return y[o];
    long long d = (l + r) / 2;
    long long s = 0;
    if (h <= d) s += q(2 * o, l, d, h, j);
    if (j > d) s += q(2 * o + 1, d + 1, r, h, j);
    return s;
}

long long f(long long o, long long l, long long r, long long h, long long j, long long v) {
    if (h > j || l > j || r < h || x[o] <= v) return -1;
    if (l == r) return l;
    long long d = (l + r) / 2;
    long long k = f(2 * o, l, d, h, j, v);
    if (k != -1) return k;
    return f(2 * o + 1, d + 1, r, h, j, v);
}

void w() {
    long long n;
    cin >> n;
    for (long long i = 1; i <= 4 * n; ++i) {
        x[i] = 0;
        y[i] = 0;
    }
    for (long long i = 1; i <= n; ++i) {
        cin >> z[i];
        g(1, 1, n, i, z[i]);
    }
    vector<long long> p(n);
    for (long long i = 0; i < n; ++i) {
        cin >> p[i];
    }
    for (long long i = 0; i < n; ++i) {
        if (i > 0) g(1, 1, n, p[i - 1], 0);
        long long t = f(1, 1, n, 1, n, 0);
        if (t == -1) {
            cout << 0 << " ";
            continue;
        }
        long long c = z[t];
        long long e = 0;
        while (true) {
            long long k = f(1, 1, n, t + 1, n, c);
            if (k == -1) break;
            long long s = q(1, 1, n, t + 1, k - 1);
            if (c + s < z[k]) {
                e++;
                c = z[k];
            } else {
                c += s + z[k];
            }
            t = k;
        }
        cout << e << " ";
    }
    cout << "\n";
}

int main() {
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    long long t;
    if (cin >> t) {
        while (t--) {
            w();
        }
    }
    return 0;
}