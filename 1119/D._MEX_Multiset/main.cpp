#include <bits/stdc++.h>
using namespace std;

void solve() {
    long long n;
    cin >> n;
    vector<long long> a(n);
    vector<long long> cnt(n + 2, 0);
    
    for(long long i = 0; i < n; ++i) {
        cin >> a[i];
        if(a[i] <= n + 1) {
            cnt[a[i]]++;
        }
    }

    long long M1 = 0, M2 = 0, M3 = 0;
    while(cnt[M1] >= 1) M1++;
    while(cnt[M2] >= 2) M2++;
    while(cnt[M3] >= 3) M3++;

    if(M1 == 0) {
        cout << "YES\n";
        cout << string(n, 'A') << "\n";
        return;
    }
    if(M2 == 0) {
        cout << "NO\n";
        return;
    }

    long long TA = (M2 + M3 >= M1) ? M1 : (M2 + M3);
    long long TB = M2;
    long long TC = M3;

    cout << "YES\n";
    string ans(n, ' ');
    vector<bool> hasA(TA, false), hasB(TB, false), hasC(TC, false);

    for(long long i = 0; i < n; ++i) {
        long long v = a[i];
        
        if(v < TA && !hasA[v]) {
            ans[i] = 'A';
            hasA[v] = true;
        } else if(v < TB && !hasB[v]) {
            ans[i] = 'B';
            hasB[v] = true;
        } else if(v < TC && !hasC[v]) {
            ans[i] = 'C';
            hasC[v] = true;
        } else {
            if(v != TA) {
                ans[i] = 'A';
            } else if(v != TB) {
                ans[i] = 'B';
            } else {
                ans[i] = 'C';
            }
        }
    }
    cout << ans << "\n";
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    long long t;
    if (cin >> t) {
        while(t--) {
            solve();
        }
    }
    return 0;
}