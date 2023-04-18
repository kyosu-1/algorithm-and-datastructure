n, x = map(int,input().split())
A = list(map(int,input().split()))

# Aの中にxがあるかどうかを判定する
if x in A:
    print("Yes")
else:
    print("No")
