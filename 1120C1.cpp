#include <bits/stdc++.h>
using namespace std;

void f() {
    long long n;
    cin >> n;
    vector<long long> d(n + 2, 0);
    for (long long i = 1; i <= n; ++i) {
        long long x;
        cin >> x;
        long long l = x * i;
        long long r = (x + 1) * i - 1;
        if (l < n) {
            d[l]++;
            d[min(r, n - 1) + 1]--;
        }
    }
    vector<long long> b;
    long long c = 0;
    for (long long i = 0; i < n; ++i) {
        c += d[i];
        if (!c) b.push_back(i);
    }
    cout << b.size() << "\n";
    for (long long v : b) cout << v << " ";
    cout << "\n";
}

int main() {
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    long long t;
    if (cin >> t) {
        while (t--) {
            f();
        }
    }
    return 0;
}