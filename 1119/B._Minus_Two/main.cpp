#include <iostream>
#include <vector>
#include <algorithm>
#include <map>

using namespace std;

void solve() {
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }

    int max_freq = 0;
    for (int step = 0; step < 40; ++step) {
        map<int, int> freq;
        for (int i = 0; i < n; ++i) {
            freq[a[i]]++;
            max_freq = max(max_freq, freq[a[i]]);
            a[i] = abs(a[i] - 2);
        }
    }
    cout << max_freq << "\n";
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