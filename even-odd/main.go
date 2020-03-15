package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, n := range nums {
		if n%2 == 0 {
			fmt.Print(i, " is even\n")
		} else {
			fmt.Print(i, " is odd\n")
		}
	}
}
