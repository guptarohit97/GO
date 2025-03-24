package main

import "fmt"

func addI(a, b int) int {
	return a + b
}
func addF(a, b float64) float64 {
	return a + b
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myNumbers interface{
	~int|~float64
}
type myAlias int

func main() {
	var n myAlias=20
	fmt.Println(addI(2,4))
	fmt.Println(addF(2.5,4.4))
	fmt.Println(addT(n,15))
	fmt.Println(addT(2.6,4.3))

}