N, X = map(int, input().split())
P = list(map(int, input().split()))

for i, val in enumerate(P):
    if val == X:
        print(i+1)
        exit()