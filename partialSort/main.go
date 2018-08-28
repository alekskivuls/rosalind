package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, n)
	for i := range nums {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	result := SortKthSmallest(nums, k-1)
	for _, num := range result {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

// SortKthSmallest sorts the slice before the kth smallest integer using quick
// select to find the kth smallest item and insertion sort to sort the
// smallest items.
func SortKthSmallest(nums []int, k int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 0
	max := len(nums)
	kthSmallestIndex := 0
	for {
		pivot := nums[r.Intn(max-min)+min]
		less := []int{}
		equal := []int{}
		greater := []int{}
		for i := min; i < max; i++ {
			if nums[i] < pivot {
				less = append(less, nums[i])
			} else if nums[i] == pivot {
				equal = append(equal, nums[i])
			} else {
				greater = append(greater, nums[i])
			}
		}
		nums = append(nums[:min], append(less, append(equal, append(greater, nums[max:]...)...)...)...)

		if min+len(less) <= k && k < max-len(greater) {
			kthSmallestIndex = k
			break
		} else if k < min+len(less) {
			max = min + len(less)
		} else {
			min = max - len(greater)
		}
	}
	// Insertion sort on k smallest slice
	for i := 0; i < kthSmallestIndex; i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return nums[:kthSmallestIndex+1]
}
