package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type UnionFind struct {
	parents []int
	ranks   []int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		ranks[i] = 1
	}
	return &UnionFind{parents, ranks}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] != x {
		uf.parents[x] = uf.Find(uf.parents[x])
	}
	return uf.parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	xRoot, yRoot := uf.Find(x), uf.Find(y)
	if xRoot == yRoot {
		return
	}
	if uf.ranks[xRoot] < uf.ranks[yRoot] {
		uf.parents[xRoot] = yRoot
	} else if uf.ranks[xRoot] > uf.ranks[yRoot] {
		uf.parents[yRoot] = xRoot
	} else {
		uf.parents[yRoot] = xRoot
		uf.ranks[xRoot]++
	}
}

func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    scanner.Scan()
    q, _ := strconv.Atoi(scanner.Text())

    uf := NewUnionFind(n)

    for i := 0; i < q; i++ {
        scanner.Scan()
        p, _ := strconv.Atoi(scanner.Text())

        scanner.Scan()
        a, _ := strconv.Atoi(scanner.Text())

        scanner.Scan()
        b, _ := strconv.Atoi(scanner.Text())

        if p == 0 {
            uf.Union(a, b)
        } else {
            if uf.IsConnected(a, b) {
                fmt.Println("Yes")
            } else {
                fmt.Println("No")
            }
        }
    }
}
