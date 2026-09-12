#include <bits/stdc++.h>
using namespace std;

struct N {
    long long s, m, fv, inv, ex;
    bool a;
} t[800005];

long long a[200005];
long long p[200005];

pair<long long, long long> g(long long o, long long l, long long r, long long v) {
    if (!t[o].a) return {0, v};
    if (v >= t[o].m) return {0, v + t[o].s};
    if (v < t[o].fv) return {t[o].inv, t[o].ex};
    
    long long d = (l + r) / 2;
    pair<long long, long long> L = g(2 * o, l, d, v);
    pair<long long, long long> R = g(2 * o + 1, d + 1, r, L.second);
    return {L.first + R.first, R.second};
}

void u(long long o, long long l, long long r) {
    long long L = 2 * o, R = 2 * o + 1;
    t[o].s = t[L].s + t[R].s;
    t[o].m = max(t[L].m, t[R].m);
    t[o].a = t[L].a | t[R].a;
    
    if (!t[o].a) {
        t[o].fv = t[o].inv = t[o].ex = 0;
        return;
    }
    
    if (t[L].a) {
        t[o].fv = t[L].fv;
        pair<long long, long long> P = g(R, (l + r) / 2 + 1, r, t[L].ex);
        t[o].inv = t[L].inv + P.first;
        t[o].ex = P.second;
    } else {
        t[o].fv = t[R].fv;
        t[o].inv = t[R].inv;
        t[o].ex = t[R].ex;
    }
}

void b(long long o, long long l, long long r) {
    if (l == r) {
        t[o].s = a[l];
        t[o].m = a[l];
        t[o].fv = a[l];
        t[o].inv = 1;
        t[o].ex = a[l];
        t[o].a = true;
        return;
    }
    long long d = (l + r) / 2;
    b(2 * o, l, d);
    b(2 * o + 1, d + 1, r);
    u(o, l, r);
}

void d(long long o, long long l, long long r, long long i) {
    if (l == r) {
        t[o].s = t[o].m = t[o].fv = t[o].inv = t[o].ex = 0;
        t[o].a = false;
        return;
    }
    long long mid = (l + r) / 2;
    if (i <= mid) d(2 * o, l, mid, i);
    else d(2 * o + 1, mid + 1, r, i);
    u(o, l, r);
}

void w() {
    long long n;
    cin >> n;
    for (long long i = 1; i <= n; i++) cin >> a[i];
    for (long long i = 1; i <= n; i++) cin >> p[i];
    
    b(1, 1, n);
    
    for (long long i = 1; i <= n; i++) {
        if (i > 1) d(1, 1, n, p[i - 1]);
        if (t[1].a) cout << t[1].inv - 1 << " ";
        else cout << 0 << " ";
    }
    cout << "\n";
}

int main() {
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    long long tc;
    if (cin >> tc) {
        while (tc--) {
            w();
        }
    }
    return 0;
}