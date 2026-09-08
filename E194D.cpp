#include <bits/stdc++.h>
using namespace std;

struct P {
    long long l, r;
};

string s;
int n;

bool chk(long long c) {
    vector<P> a;
    a.push_back({0, 0});
    for (int i = 0; i < n; i++) {
        vector<P> b;
        for (auto x : a) {
            if (x.l < x.r) {
                b.push_back({x.l - c, x.r + c});
            } else {
                long long v = x.l;
                b.push_back({v - c, v - 1});
                b.push_back({v + 1, v + c});
            }
        }
        vector<P> d;
        char ch = s[i];
        for (auto x : b) {
            long long nl = x.l, nr = x.r;
            if (ch == '+') {
                nl = max(nl, 1LL);
            } else if (ch == '-') {
                nr = min(nr, -1LL);
            } else {
                nl = max(nl, 0LL);
                nr = min(nr, 0LL);
            }
            if (nl <= nr) {
                d.push_back({nl, nr});
            }
        }
        if (d.empty()) return false;
        
        vector<P> m;
        for (auto x : d) {
            if (m.empty()) {
                m.push_back(x);
            } else {
                if (x.l <= m.back().r + 1) {
                    m.back().r = max(m.back().r, x.r);
                } else {
                    m.push_back(x);
                }
            }
        }
        a = m;
    }
    return !a.empty();
}

void run() {
    cin >> n >> s;
    if (s[0] == '0') {
        cout << -1 << "\n";
        return;
    }
    for (int i = 1; i < n; i++) {
        if (s[i] == '0' && s[i - 1] == '0') {
            cout << -1 << "\n";
            return;
        }
    }
    long long lo = 1, hi = 2LL * n + 5, ans = -1;
    while (lo <= hi) {
        long long mid = lo + (hi - lo) / 2;
        if (chk(mid)) {
            ans = mid;
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }
    cout << ans << "\n";
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    if (cin >> t) {
        while (t--) {
            run();
        }
    }
    return 0;
}