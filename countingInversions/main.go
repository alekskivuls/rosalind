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
	_, numInversions := InsertionSort(nums)
	fmt.Println(numInversions)
}

// InsertionSort sorts the given int slice using insertion sort and counts the
//  number of inversions it took to sort
func InsertionSort(nums []int) (sorted []int, numInversions int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
			numInversions++
		}
	}
	return nums, numInversions
}
