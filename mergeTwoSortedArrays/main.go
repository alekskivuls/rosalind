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
	b := readIntSlice(scanner)
	c := merge(a, b)
	for _, num := range c {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func readIntSlice(scanner *bufio.Scanner) []int {
	slice := make([]int, 0)
	scanner.Scan()
	sliceLen, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < sliceLen; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		slice = append(slice, num)
	}
	return slice
}

func merge(left []int, right []int) []int {
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
