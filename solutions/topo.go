package main

import (
	"container/list"
	"fmt"
)

// Function to perform Topological Sort using Kahn's Algorithm
func topologicalSort(vertices int, edges [][]int) ([]int, error) {
	// Create an adjacency list to represent the graph
	graph := make([][]int, vertices)
	inDegree := make([]int, vertices) // In-degree for each vertex

	// Build the graph and calculate in-degrees
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		inDegree[v]++
	}

	// Initialize a queue for vertices with in-degree of 0
	queue := list.New()
	for i := 0; i < vertices; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	// List to store the topological sort
	var topOrder []int

	// Perform Kahn's Algorithm
	for queue.Len() > 0 {
		// Dequeue a vertex with in-degree 0
		node := queue.Front().Value.(int)
		queue.Remove(queue.Front())

		// Add it to the topological order
		topOrder = append(topOrder, node)

		// For each neighbor of this node, reduce its in-degree
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			// If in-degree becomes 0, add it to the queue
			if inDegree[neighbor] == 0 {
				queue.PushBack(neighbor)
			}
		}
	}

	// If the topological order doesn't contain all vertices, there's a cycle
	if len(topOrder) != vertices {
		return nil, fmt.Errorf("the graph contains a cycle, topological sort not possible")
	}

	return topOrder, nil
}

func runTopo() {
	// Example Graph (represented as edges)
	// Graph: A -> B -> D
	//         ↓    ↑
	//         C -> E
	edges := [][]int{
		{0, 1}, // A -> B
		{1, 3}, // B -> D
		{0, 2}, // A -> C
		{2, 4}, // C -> E
		{4, 3}, // E -> D
	}

	// Number of vertices in the graph (A=0, B=1, C=2, D=3, E=4)
	vertices := 5

	// Perform topological sort
	result, err := topologicalSort(vertices, edges)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Topological Sort:", result)
	}
}
