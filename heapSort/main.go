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
	for _, num := range HeapSort(nums) {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

// MaxHeap stores the data for a max heap
type MaxHeap struct {
	heap []int
}

// NewMaxHeap creates a new max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{make([]int, 0)}
}

// Push pushes data onto the max heap
func (h *MaxHeap) Push(num int) {
	pos := len(h.heap)
	h.heap = append(h.heap, num)
	for (pos+1)/2-1 >= 0 && h.heap[pos] > h.heap[(pos+1)/2-1] {
		h.heap[(pos+1)/2-1], h.heap[pos] = h.heap[pos], h.heap[(pos+1)/2-1]
		pos = (pos+1)/2 - 1
	}
}

// Pop pops off data from the max heap
func (h *MaxHeap) Pop() int {
	max := h.heap[0]
	h.heap[0], h.heap[len(h.heap)-1] = h.heap[len(h.heap)-1], h.heap[0]
	h.heap = h.heap[0 : len(h.heap)-1]
	pos := 0
	for {
		left := 2*(pos+1) - 1
		right := 2 * (pos + 1)
		switchIndex := left
		if right < len(h.heap) && h.heap[right] > h.heap[left] {
			switchIndex = right
		}
		if switchIndex < len(h.heap) && h.heap[pos] < h.heap[switchIndex] {
			h.heap[pos], h.heap[switchIndex] = h.heap[switchIndex], h.heap[pos]
			pos = switchIndex
			continue
		}
		break
	}
	return max
}

// Len returns the size of the heap
func (h *MaxHeap) Len() int {
	return len(h.heap)
}

// HeapSort takes a slice of ints and sorts them using heap sort
func HeapSort(nums []int) []int {
	heap := NewMaxHeap()
	for _, num := range nums {
		heap.Push(num)
	}
	sorted := make([]int, len(nums))
	for i := 1; i <= len(nums); i++ {
		sorted[len(nums)-i] = heap.Pop()
	}
	return sorted
}
