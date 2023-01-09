package graph

import (
	"math"

	"github.com/kyosu-1/algorithm-and-datastructure/datastructure/heapq"
)

// Dijkstra performs a Dijkstra's algorithm on a graph starting at vertex
// the given source vertex, and returns the shortest path to each vertex.
// use priority queue data structure
func Dijkstra(g WeightedGraph, s int) []int {
	visited := make([]bool, g.Order())
	dist := make([]int, g.Order())
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[s] = 0
	q := heapq.New[int]()
	d := heapq.NewData(1, s)
	q.Push(d)
	for !q.IsEmpty() {
		v := q.Pop()
		visited[v.Value] = true
		for _, w := range g.Adjacent(v.Value) {
			if !visited[w] {
				wegiht := g.Weight(v.Value, w)
				if dist[v.Value] + wegiht < dist[w] {
					dist[w] = dist[v.Value] + wegiht
					d := heapq.NewData(dist[w], w)
					q.Push(d)
				}
			}
		}
	}
	return dist
}
