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
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for arrNum := 0; arrNum < k; arrNum++ {
		arr := make([]int, 0)
		for i := 0; i < n; i++ {
			scanner.Scan()
			k, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			arr = append(arr, k)
		}
		if arrNum != k-1 {
			fmt.Printf("%d ", majorityElement(arr))
		} else {
			fmt.Printf("%d\n", majorityElement(arr))
		}
	}

}

func majorityElement(arr []int) int {
	majorElement := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == majorElement {
			count++
		} else {
			count--
			if count == 0 {
				majorElement = arr[i]
				count++
			}
		}
	}

	count = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == majorElement {
			count++
		}
	}
	if count > len(arr)/2 {
		return majorElement
	}
	return -1
}
