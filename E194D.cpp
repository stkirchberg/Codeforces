#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long n;
    cin >> n;
    string s;
    cin >> s;
    
    long long m = 0;
    long long c = 0;
    
    for (long long i = 0; i < n; i++) {
        if (i == 0) {
            c = 1;
        } else {
            if (s[i] == s[i - 1]) {
                c++;
            } else {
                c = 1;
            }
        }
        m = max(m, c);
    }
    
    cout << m << "\n";
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