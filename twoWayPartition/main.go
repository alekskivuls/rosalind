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

	for _, num := range TwoWayPartition(nums) {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

//TwoWayPartition takes an int slice and permutes the slice so that all
//elemnets left of the pivot are less than or equal and all elements to the
//right are greater than the pivot. Takes the first element as the pivot.
func TwoWayPartition(nums []int) []int {
	pivot := nums[0]
	left := make([]int, 0)
	right := make([]int, 0)
	right = append(right, pivot)

	for _, val := range nums[1:len(nums)] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return append(left, right...)
}
