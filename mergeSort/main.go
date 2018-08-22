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
	a := readIntSlice(scanner)
	sorted := MergeSort(a)
	for _, num := range sorted {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func readIntSlice(scanner *bufio.Scanner) []int {
	scanner.Scan()
	sliceLen, _ := strconv.Atoi(scanner.Text())
	slice := make([]int, sliceLen)
	for i := 0; i < sliceLen; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		slice[i] = num
	}
	return slice
}

//MergeSort sorts a int slice and returns the sorted result
func MergeSort(slice []int) []int {
	return merge(slice[0:len(slice)/2], slice[len(slice)/2:len(slice)])
}

func merge(left []int, right []int) []int {
	if len(left) > 1 {
		left = merge(left[0:len(left)/2], left[len(left)/2:len(left)])
	}
	if len(right) > 1 {
		right = merge(right[0:len(right)/2], right[len(right)/2:len(right)])
	}
	merged := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else if left[i] > right[j] {
			merged = append(merged, right[j])
			j++
		} else {
			merged = append(merged, left[i])
			i++
			merged = append(merged, right[j])
			j++
		}
	}
	if i < len(left) {
		for i < len(left) {
			merged = append(merged, left[i])
			i++
		}
	}
	if j < len(right) {
		for j < len(right) {
			merged = append(merged, right[j])
			j++
		}
	}
	return merged
}
