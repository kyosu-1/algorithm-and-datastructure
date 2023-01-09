package graph

import (
	"github.com/kyosu-1/algorithm-and-datastructure/datastructure/queue"
)

// BreadthFirstSearch performs a breadth-first search of the graph g, starting at
// the given source vertex, and calls the function f for each vertex visited.
// use queue data structure
func BreadthFirstSearch(g Graph, s int, f func(int)) {
	visited := make([]bool, g.Order())
	visited[s] = true
	q := queue.New[int]()
	q.Enqueue(s)
	for !q.IsEmpty() {
		v := q.Dequeue()
		f(v)
		for _, w := range g.Adjacent(v) {
			if !visited[w] {
				visited[w] = true
				q.Enqueue(w)
			}
		}
	}
}
