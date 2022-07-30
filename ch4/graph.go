package ch4

func addEdge(from, to string, graph map[string]map[string]bool) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string, graph map[string]map[string]bool) bool {
	return graph[from][to]
}
