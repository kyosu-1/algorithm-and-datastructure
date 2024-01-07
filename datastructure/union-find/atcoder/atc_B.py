class UnionFind:

    def __init__(self, n):
        # -1で初期化する
        self.parents = [-1] * n
        # 木の高さを格納する
        self.rank = [0] * n
    
    # xの根を探す
    def find(self, x):
        if self.parents[x] == -1:
            return x
        else:
            # 経路圧縮
            self.parents[x] = self.find(self.parents[x])
            return self.parents[x]
    
    # xとyの属する集合を併合する
    def unite(self, x, y):
        # x, yの根を探す
        x = self.find(x)
        y = self.find(y)

        # 同じ集合に属している場合
        if x == y:
            return
        
        # 高さが低い方を高い方に併合する
        if self.rank[x] < self.rank[y]:
            self.parents[x] = y
        else:
            self.parents[y] = x
            # 高さが同じ場合は高さを1増やす
            if self.rank[x] == self.rank[y]:
                self.rank[x] += 1
    
    # xとyが同じ集合に属するか判定する
    def same(self, x, y):
        return self.find(x) == self.find(y)


if __name__ == '__main__':
    n, q = map(int, input().split())
    uf = UnionFind(n)
    for _ in range(q):
        p, a, b = map(int, input().split())
        if p == 0:
            uf.unite(a, b)
        else:
            if uf.same(a, b):
                print('Yes')
            else:
                print('No')
