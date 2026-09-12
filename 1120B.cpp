#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long n, k;
    cin >> n >> k;
    
    if (k < n || k >= 2 * n) {
        cout << -1 << "\n";
        return;
    }
    
    long long c = 2 * n - k;
    vector<vector<long long>> a(n, vector<long long>(n, 0));
    long long v = 1;
    
    for (long long i = 0; i < c; ++i) {
        a[i][i] = v++;
    }
    
    for (long long i = c; i < n; ++i) {
        a[i][0] = v++;
    }
    
    for (long long j = c; j < n; ++j) {
        a[0][j] = v++;
    }
    
    for (long long i = 0; i < n; ++i) {
        for (long long j = 0; j < n; ++j) {
            if (a[i][j] == 0) {
                a[i][j] = v++;
            }
        }
    }
    
    for (long long i = 0; i < n; ++i) {
        for (long long j = 0; j < n; ++j) {
            cout << a[i][j] << (j == n - 1 ? "" : " ");
        }
        cout << "\n";
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    
    long long t;
    if (cin >> t) {
        while (t--) {
            solve();
        }
    }
    
    return 0;
}