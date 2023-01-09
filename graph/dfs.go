package graph

import (
	"github.com/kyosu-1/algorithm-and-datastructure/datastructure/stack"
)

// DepthFirstSearch performs a depth first search on a graph starting at vertex
// the given source vertex, and calls the function f for each vertex visited.
// use stack data structure
func DepthFirstSearch(g Graph, s int, f func(int)) {
	visited := make([]bool, g.Order())
	visited[s] = true
	q := stack.New[int]()
	q.Push(s)
	for !q.IsEmpty() {
		v := q.Pop()
		f(v)
		for _, w := range g.Adjacent(v) {
			if !visited[w] {
				visited[w] = true
				q.Push(w)
			}
		}
	}
}
