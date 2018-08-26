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
	fmt.Println(FindKthSmallest(nums, k-1))
}

// FindKthSmallest finds the kth smallest integer in a slice using quick select
func FindKthSmallest(nums []int, k int) (kthSmallest int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 0
	max := len(nums)
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
			kthSmallest = equal[0]
			break
		} else if k < min+len(less) {
			max = min + len(less)
		} else {
			min = max - len(greater)
		}
	}
	return kthSmallest
}
