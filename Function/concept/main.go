package main

import "fmt"

func main() {
	x := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func sum(a []int) (total int) {
	total = 0

	for _, v := range a {
		total += v
	}
	return total
}