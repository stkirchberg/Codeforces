#include <bits/stdc++.h>
using namespace std;

void solve() {
    int n;
    cin >> n;
    vector<int> a(n);
    vector<int> p;
    int first = -1, last = -1;

    for (int i = 0; i < n; i++) {
        cin >> a[i];
        if (a[i] != 0) {
            if (first == -1) first = i;
            last = i;
        }
        if (a[i] == 1) {
            p.push_back(i);
        }
    }

    if (p.empty()) {
        if (first != -1) {
            a[first] = 1;
            a[last] = 1;
        }
    } else {
        int best_len = 0;
        int type = 0;

        for (int i = 0; i + 1 < (int)p.size(); i++) {
            int len = p[i + 1] - p[i] + 1;
            if (len > best_len) {
                best_len = len;
                type = 0;
            }
        }
        if (p[0] - first + 1 > best_len) {
            best_len = p[0] - first + 1;
            type = 1;
        }
        if (last - p.back() + 1 > best_len) {
            best_len = last - p.back() + 1;
            type = 2;
        }

        if (type == 1) a[first] = 1;
        if (type == 2) a[last] = 1;
    }

    for (int i = 0; i < n; i++) {
        if (a[i] == -1) a[i] = 0;
        cout << a[i] << (i == n - 1 ? "" : " ");
    }
    cout << "\n";
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