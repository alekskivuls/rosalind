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
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	for _, num := range BuildMaxHeap(nums) {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

//BuildMaxHeap returns a slice of ints with the ordering as if they had been
//inserted into a max heap
func BuildMaxHeap(nums []int) []int {
	//Make heap indexing one based for easier math
	heap := make([]int, 1)
	for _, num := range nums {
		pos := len(heap)
		heap = append(heap, num)
		for pos/2 > 0 && heap[pos] > heap[pos/2] {
			heap[pos/2], heap[pos] = heap[pos], heap[pos/2]
			pos = pos / 2
		}
	}
	return heap[1:len(heap)]
}
