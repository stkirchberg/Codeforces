#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long a, b;
    cin >> a >> b;
    long long s = a + b;
    long long r = 0;
    
    for (long long i = 29; i >= 0; i--) {
        if (s & (1LL << i)) {
            if ((r | (1LL << i)) <= a) {
                r |= (1LL << i);
            }
        }
    }
    
    long long o = a - r;
    cout << s << " " << o << "\n";
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