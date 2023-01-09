package graph

type AdjacencyList map[int][]int

// NewAdjacencyList returns a new adjacency list
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{}
}

// AddEdge adds an edge to the adjacency list
func (a *AdjacencyList) AddEdge(v, w int) {
	(*a)[v] = append((*a)[v], w)
}

// RemoveEdge removes an edge from the adjacency list
func (a *AdjacencyList) RemoveEdge(v, w int) {
	adjacent := (*a)[v]
	for i, vertex := range adjacent {
		if vertex == w {
			(*a)[v] = append(adjacent[:i], adjacent[i+1:]...)
			break
		}
	}
}

// HasEdge returns true if there is an edge between v and w
func (a *AdjacencyList) HasEdge(v, w int) bool {
	adjacent := (*a)[v]
	for _, vertex := range adjacent {
		if vertex == w {
			return true
		}
	}
	return false
}

// Adjacent returns the vertices adjacent to v
func (a *AdjacencyList) Adjacent(v int) []int {
	return (*a)[v]
}

// Order returns the number of vertices in the graph
func (a *AdjacencyList) Order() int {
	return len(*a)
}

// InDegree returns the in-degree of vertex v
func (a *AdjacencyList) InDegree(v int) int {
	degree := 0
	for _, adjacent := range *a {
		for _, vertex := range adjacent {
			if vertex == v {
				degree++
			}
		}
	}
	return degree
}

// OutDegree returns the out-degree of vertex v
func (a *AdjacencyList) OutDegree(v int) int {
	return len((*a)[v])
}

// WeightedAdjacencyList is a weighted adjacency list
type WeightedAdjacencyList map[int]map[int]int

// NewWeightedAdjacencyList returns a new weighted adjacency list
func NewWeightedAdjacencyList() *WeightedAdjacencyList {
	return &WeightedAdjacencyList{}
}

// AddEdge adds an edge to the weighted adjacency list
func (a *WeightedAdjacencyList) AddEdge(v, w, weight int) {
	if (*a)[v] == nil {
		(*a)[v] = make(map[int]int)
	}
	(*a)[v][w] = weight
}

// RemoveEdge removes an edge from the weighted adjacency list
func (a *WeightedAdjacencyList) RemoveEdge(v, w int) {
	delete((*a)[v], w)
}

// HasEdge returns true if there is an edge between v and w
func (a *WeightedAdjacencyList) HasEdge(v, w int) bool {
	_, ok := (*a)[v][w]
	return ok
}

// Adjacent returns the vertices adjacent to v
func (a *WeightedAdjacencyList) Adjacent(v int) []int {
	adjacent := make([]int, 0, len((*a)[v]))
	for vertex := range (*a)[v] {
		adjacent = append(adjacent, vertex)
	}
	return adjacent
}

// Order returns the number of vertices in the graph
func (a *WeightedAdjacencyList) Order() int {
	return len(*a)
}

// InDegree returns the in-degree of vertex v
func (a *WeightedAdjacencyList) InDegree(v int) int {
	degree := 0
	for _, adjacent := range *a {
		for vertex := range adjacent {
			if vertex == v {
				degree++
			}
		}
	}
	return degree
}

// OutDegree returns the out-degree of vertex v
func (a *WeightedAdjacencyList) OutDegree(v int) int {
	return len((*a)[v])
}

// Weight returns the weight of the edge between vertices v and w
func (a *WeightedAdjacencyList) Weight(v, w int) int {
	return (*a)[v][w]
}
