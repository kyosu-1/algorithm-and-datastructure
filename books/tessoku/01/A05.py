N, K = map(int, input().split())

ans = 0

for i in range(1, N + 1):
    for k in range(1, N + 1):
        c = K - i - k
        if 1 <= c <= N:
            ans += 1

print(ans)
