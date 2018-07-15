package main

import "fmt"

func main() {
	n := 100000
	m := 100000
	fmt.Println(n)
	fmt.Println(m)
	fmt.Print(n)
	for i := 0; i < n; i++ {
		fmt.Printf(" %d", n)
	}
	fmt.Println()
	fmt.Print(n)
	for i := 0; i < n; i++ {
		fmt.Printf(" %d", n)
	}
}
