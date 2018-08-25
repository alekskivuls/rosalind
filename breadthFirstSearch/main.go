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

	adjList := make(map[int][]int, 0)
	for i := 1; i <= numNodes; i++ {
		adjList[i] = make([]int, 0)
	}
	for i := 1; i <= numEdges; i++ {
		scanner.Scan()
		left, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		right, _ := strconv.Atoi(scanner.Text())
		adjList[left] = append(adjList[left], right)
	}
	distances := FindDistances(adjList, 1)
	for i := 1; i <= numNodes; i++ {
		if distance, ok := distances[i]; ok {
			fmt.Print(distance, " ")
		} else {
			fmt.Print(-1, " ")
		}
	}
	fmt.Println()
}

// FindDistances finds all the nodes connected to the start node and returns
// their distance from the start node in a map
func FindDistances(adjList map[int][]int, start int) map[int]int {
	// Check if the start node exists in the adjacency list
	if _, ok := adjList[start]; !ok {
		return nil
	}
	// BFS
	visited := make(map[int]int, 0)
	visited[start] = 0
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) != 0 {
		// Pop node off queue
		node := queue[0]
		queue = queue[1:len(queue)]
		// Check all adjacent nodes if they are visited yet
		for _, adjNode := range adjList[node] {
			if _, ok := visited[adjNode]; !ok {
				visited[adjNode] = visited[node] + 1
				queue = append(queue, adjNode)
			}
		}
	}
	return visited
}
