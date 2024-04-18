package algorithms

import "math"

type Edge struct {
	from, to, cost int
}

func BellmanFord(edges []Edge, vertices, source int) ([]float64, bool) {
	distance := make([]float64, vertices)
	// Step 1: Initialize distances for all vertices to infinity
	for i := range distance {
		distance[i] = math.Inf(1)
	}
	distance[source] = 0

	// Relax all edges V - 1 times
	for i := 0; i < vertices-1; i++ {
		for _, edge := range edges {
			if distance[edge.from]+float64(edge.cost) < distance[edge.to] {
				distance[edge.to] = distance[edge.from] + float64(edge.cost)
			}
		}
	}

	// Check for negative weight cycles
	for _, edge := range edges {
		if distance[edge.from]+float64(edge.cost) < distance[edge.to] {
			return distance, true
		}
	}
	return distance, false
}
