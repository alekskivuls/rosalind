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
	numNodes, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	adjList := make(map[int][]int, 0)
	for i := 1; i <= numNodes; i++ {
		adjList[i] = make([]int, 0)
	}

	scanner.Scan()
	numEdges, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < numEdges; i++ {
		scanner.Scan()
		startEdge, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		endEdge, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		adjList[startEdge] = append(adjList[startEdge], endEdge)
		adjList[endEdge] = append(adjList[endEdge], startEdge)
	}

	for i := 1; i <= numNodes; i++ {
		if i != 1 {
			fmt.Print(" ")
		}
		fmt.Print(len(adjList[i]))
	}
	fmt.Println()
}
