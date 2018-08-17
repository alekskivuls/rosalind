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

	adjList := make(map[int][]int, 0)
	for i := 1; i <= numNodes; i++ {
		adjList[i] = make([]int, 0)
	}

	scanner.Scan()
	numEdges, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < numEdges; i++ {
		scanner.Scan()
		startEdge, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		endEdge, _ := strconv.Atoi(scanner.Text())
		adjList[startEdge] = append(adjList[startEdge], endEdge)
		adjList[endEdge] = append(adjList[endEdge], startEdge)
	}

	for i := 1; i <= numNodes; i++ {
		count := 0
		for j := 0; j < len(adjList[i]); j++ {
			count += len(adjList[adjList[i][j]])
		}
		if i != 1 {
			fmt.Print(" ")
		}
		fmt.Print(count)
	}
	fmt.Println()
}
