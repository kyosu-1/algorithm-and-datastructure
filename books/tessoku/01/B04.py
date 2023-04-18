N = str(input())

ans = 0

# Nは2進法で表されており、それを10進法に変換する
for i in range(len(N)):
    ans += int(N[i]) * (1 << (len(N) - i - 1))

print(ans)