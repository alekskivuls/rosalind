package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numNodes, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	numEdges, _ := strconv.Atoi(scanner.Text())

	adjList := make(map[int][]int, numNodes)
	for i := 1; i <= numNodes; i++ {
		adjList[i] = make([]int, 0)
	}

	for i := 1; i <= numEdges; i++ {
		scanner.Scan()
		left, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		right, _ := strconv.Atoi(scanner.Text())
		adjList[left] = append(adjList[left], right)
		adjList[right] = append(adjList[right], left)
	}
	scc := GetStronglyConnectedComponents(adjList)
	fmt.Println(len(scc))
}

// GetStronglyConnectedComponents gets the strongly connected componenets
// given an adjacency list and returns the results as a slice of slices of each
// component with the nodes that are contained in that component
func GetStronglyConnectedComponents(adjList map[int][]int) [][]int {
	scc := make([][]int, 0)
	visited := make(map[int]struct{}, 0)
	var s struct{}
	// DFS from each node
	for node := range adjList {
		// Check if we already visited the node, if we have skip it
		if _, ok := visited[node]; ok {
			continue
		}
		// Mark all visited nodes in the overall visited and add the set to scc
		component := dfs(adjList, node)
		for _, visitedNode := range component {
			visited[visitedNode] = s
		}
		scc = append(scc, component)
	}
	return scc
}

func dfs(adjList map[int][]int, start int) []int {
	stack := make([]int, 0)
	stack = append(stack, start)

	var s struct{}
	visited := make(map[int]struct{}, 0)
	visited[start] = s

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		found := false
		for _, adjNode := range adjList[node] {
			if _, ok := visited[adjNode]; !ok {
				stack = append(stack, adjNode)
				visited[adjNode] = s
				found = true
				break
			}
		}
		// If no unvisited nodes are adjacent pop the stack
		if !found {
			stack = stack[0 : len(stack)-1]
		}
	}

	result := make([]int, 0)
	for node := range visited {
		result = append(result, node)
	}
	return result
}
