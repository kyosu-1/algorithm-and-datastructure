N = int(input())
A = list(map(int, input().split()))

# 累積和
S = [0] * (N+1)
for i in range(N):
    S[i+1] = S[i] + A[i]

Q = int(input())
for _ in range(Q):
    L, R = map(int, input().split())
    diff = (R - L + 1) - 2 * (S[R] - S[L-1])
    if diff > 0:
        print("lose")
    elif diff == 0:
        print("draw")
    else:
        print("win")