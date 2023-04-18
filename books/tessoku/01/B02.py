a, b = map(int, input().split())

if len([i for i in range(a, b + 1) if 100 % i == 0]) > 0:
    print("Yes")
else:
    print("No")
