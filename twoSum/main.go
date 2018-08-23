package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type (
	//TwoSum represents the indicies of a two sum
	TwoSum struct {
		left, right int
	}
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
		twoSum, found := FindTwoSum(nums)
		if !found {
			fmt.Println(-1)
		} else {
			fmt.Println(twoSum.left+1, twoSum.right+1)
		}
	}
}

//FindTwoSum returns a twoSum if found otherwise returns false
func FindTwoSum(nums []int) (twoSum TwoSum, found bool) {
	hashMap := make(map[int]int, 0)
	for rightIndex, val := range nums {
		if leftIndex, ok := hashMap[-val]; ok {
			return TwoSum{leftIndex, rightIndex}, true
		}
		hashMap[val] = rightIndex
	}
	return twoSum, false
}
