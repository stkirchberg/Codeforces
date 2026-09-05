#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long n;
    cin >> n;
    vector<long long> b(n);
    vector<long long> ans(n, -1);
    bool ok = true;

    for(long long i = 0; i < n; i++) {
        cin >> b[i];
    }

    for(long long i = 0; i < n; i++) {
        if(b[i] == 0) {
            if(ans[i] == 0) ok = false;
            ans[i] = 1;
        } else if(b[i] > 0) {
            long long d = b[i];
            long long l = max(0LL, i - d);
            long long r = min(n - 1, i + d);
            for(long long j = l; j <= r; j++) {
                if(j == i - d || j == i + d) continue;
                if(ans[j] == 1) ok = false;
                ans[j] = 0;
            }
            long long left_pos = i - d;
            long long right_pos = i + d;
            bool can_l = (left_pos >= 0 && ans[left_pos] != 0);
            bool can_r = (right_pos < n && ans[right_pos] != 0);
            if(left_pos < 0 && right_pos >= n) ok = false;
            else if(left_pos < 0) {
                if(ans[right_pos] == 0) ok = false;
                ans[right_pos] = 1;
            } else if(right_pos >= n) {
                if(ans[left_pos] == 0) ok = false;
                ans[left_pos] = 1;
            } else {
                if(!can_l && !can_r) ok = false;
                else if(can_l && !can_r) ans[left_pos] = 1;
                else if(!can_l && can_r) ans[right_pos] = 1;
            }
        }
    }

    if(!ok) {
        cout << -1 << "\n";
        return;
    }

    long long count1 = 0;
    for(long long i = 0; i < n; i++) {
        if(ans[i] == 1) count1++;
    }
    if(count1 == 0) {
        bool placed = false;
        for(long long i = 0; i < n; i++) {
            if(ans[i] == -1) {
                ans[i] = 1;
                placed = true;
                break;
            }
        }
        if(!placed) {
            for(long long i = 0; i < n; i++) {
                if(ans[i] == 0) {
                    ans[i] = 1;
                    placed = true;
                    break;
                }
            }
        }
        if(!placed) {
            cout << -1 << "\n";
            return;
        }
    }

    for(long long i = 0; i < n; i++) {
        cout << (ans[i] == 1 ? 1 : 0);
    }
    cout << "\n";
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    long long t;
    if(cin >> t) {
        while(t--) {
            solve();
        }
    }
    return 0;
}