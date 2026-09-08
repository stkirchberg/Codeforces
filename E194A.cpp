#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long n;
    cin >> n;
    vector<long long> v(n);
    long long z = 0;
    
    for (long long i = 0; i < n; i++) {
        cin >> v[i];
        if (v[i] == 0) {
            z++;
        }
    }

    if (v[0] == 0 && v[n - 1] == 0) {
        cout << 0 << "\n";
    } else if (v[0] == 1 && v[n - 1] == 1) {
        if (z >= 2) {
            cout << 2 << "\n";
        } else {
            cout << -1 << "\n";
        }
    } else {
        if (z >= 2) {
            cout << 1 << "\n";
        } else {
            cout << -1 << "\n";
        }
    }
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