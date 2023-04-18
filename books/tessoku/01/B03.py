N = int(input())
A = list(map(int, input().split()))

# Aから異なる3つの要素を選び、その和が1000になるものがあるかどうかを判定する
# 1000になるものがあるならYes、ないならNoを出力する
def is_exist(A):
    for i in range(len(A)):
        for j in range(i + 1, len(A)):
            for k in range(j + 1, len(A)):
                if A[i] + A[j] + A[k] == 1000:
                    return True
    return False

print("Yes" if is_exist(A) else "No")
