N, X, Y = map(int, input().split())

a = [0 for i in range(N)]
b = [0 for i in range(N)]
a[N-1] = 1

for i in range(N-2, -1, -1):
    a[i] += a[i+1]
    b[i+1] += X*a[i+1]
    a[i] += b[i+1]
    b[i] += Y*b[i+1]

print(b[0])
