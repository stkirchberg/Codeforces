#include <bits/stdc++.h>
using namespace std;

struct P {
    long long l, r;
};

string s;
int n;

bool chk(int c) {
    P a[4];
    int na = 1;
    a[0] = {0, 0};
    
    for (int i = 0; i < n; i++) {
        P b[8];
        int nb = 0;
        for (int j = 0; j < na; j++) {
            if (a[j].l < a[j].r) {
                b[nb++] = {a[j].l - c, a[j].r + c};
            } else {
                long long v = a[j].l;
                b[nb++] = {v - c, v - 1};
                b[nb++] = {v + 1, v + c};
            }
        }
        
        P d[8];
        int nd = 0;
        char ch = s[i];
        for (int j = 0; j < nb; j++) {
            long long nl = b[j].l, nr = b[j].r;
            if (ch == '+') {
                nl = max(nl, 1LL);
            } else if (ch == '-') {
                nr = min(nr, -1LL);
            } else {
                nl = max(nl, 0LL);
                nr = min(nr, 0LL);
            }
            if (nl <= nr) {
                d[nd++] = {nl, nr};
            }
        }
        
        if (nd == 0) return false;
        
        for (int j = 0; j < nd; j++) {
            for (int k = j + 1; k < nd; k++) {
                if (d[j].l > d[k].l) {
                    swap(d[j], d[k]);
                }
            }
        }
        
        na = 0;
        for (int j = 0; j < nd; j++) {
            if (na == 0) {
                a[na++] = d[j];
            } else {
                if (d[j].l <= a[na - 1].r + 1) {
                    a[na - 1].r = max(a[na - 1].r, d[j].r);
                } else {
                    a[na++] = d[j];
                }
            }
        }
    }
    return na > 0;
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
    if (chk(1)) {
        cout << 1 << "\n";
    } else if (chk(2)) {
        cout << 2 << "\n";
    } else {
        cout << 3 << "\n";
    }
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