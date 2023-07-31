#include <iostream>
using namespace std;

int N, A[100009], S[100009];
int Q, L[100009], R[100009];

int main() {
    // 入力
    cin >> N >> Q;
    for (int i = 0; i < N; ++i) cin >> A[i];
    for (int i = 0; i < Q; ++i) cin >> L[i] >> R[i];

    // 累積和
    S[0] = 0;
    for (int i = 0; i < N; ++i) S[i + 1] = S[i] + A[i];

    // クエリに答える
    for (int i = 0; i < Q; ++i) cout << S[R[i]] - S[L[i]-1] << endl;

    return 0;
}