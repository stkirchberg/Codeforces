#include <bits/stdc++.h>
using namespace std;

void solve() {
    int n;
    cin >> n;
    int odd = 0, mod0 = 0, mod2 = 0;
    for (int i = 0; i < n; i++) {
        int x;
        cin >> x;
        if (x & 1) {
            odd++;
        } else if (x % 4 == 0) {
            mod0++;
        } else {
            mod2++;
        }
    }
    cout << max({odd, mod0, mod2}) << "\n";
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while (t--) {
        solve();
    }
    return 0;
}