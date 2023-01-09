package graph

// Graph is an interface for a graph data structure.
type Graph interface {
	// Order returns the number of vertices in the graph.
	Order() int

	// Adjacent returns the vertices adjacent to v.
	Adjacent(v int) []int

	// AddEdge adds an edge between vertices v and w.
	AddEdge(v, w int)

	// RemoveEdge removes the edge between vertices v and w.
	RemoveEdge(v, w int)

	// HasEdge returns true if there is an edge between vertices v and w.
	HasEdge(v, w int) bool
}

// WeightedGraph is an interface for a weighted graph data structure.
type WeightedGraph interface {
	Graph
	// Weight returns the weight of the edge between vertices v and w.
	Weight(v, w int) int
}

// DirectedGraph is an interface for a directed graph data structure.
type DirectedGraph interface {
	Graph
	// InDegree returns the in-degree of vertex v.
	InDegree(v int) int

	// OutDegree returns the out-degree of vertex v.
	OutDegree(v int) int
}
