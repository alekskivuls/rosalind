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
	arrSize, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, arrSize)
	for i := 0; i < len(arr); i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	_, numSwaps := insertionSort(arr)
	fmt.Printf("%d\n", numSwaps)
}

func insertionSort(arr []int) (sorted []int, numSwaps int) {
	count := 0
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			count++
		}
	}
	return arr, count
}
