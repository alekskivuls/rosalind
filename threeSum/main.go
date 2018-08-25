package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for caseNum := 0; caseNum < k; caseNum++ {
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			nums[i], _ = strconv.Atoi(scanner.Text())
		}
		if result := GetThreeSum(nums); result != nil {
			fmt.Println(result[0]+1, result[1]+1, result[2]+1)
		} else {
			fmt.Println(-1)
		}
	}
}

// GetThreeSum finds a three sum and returns the indicies in the result.
// Returns nil if a three sum is not found.
func GetThreeSum(nums []int) []int {
	indexMap := make(map[int]int, len(nums))
	for index, num := range nums {
		indexMap[-num] = index
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if index, ok := indexMap[nums[i]+nums[j]]; ok {
				indices := []int{index, i, j}
				sort.Ints(indices)
				return indices
			}
		}
	}
	return nil
}
