def make_number(n):
    return int('1' * n)

# 上記のような数で、N番目に小さい数を求める
# 1 <= N <= 333
def solve():
    repunt_set = set()
    for i in range(1, 100):
        for j in range(1, 100):
            for k in range(1, 100):
                repunt_set.add(make_number(i) + make_number(j) + make_number(k))
    
    repunts = list(repunt_set)
    repunts.sort()

    N = int(input())
    print(repunts[N-1])

if __name__ == '__main__':
    solve()
