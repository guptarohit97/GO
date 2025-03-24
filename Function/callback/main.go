package main

import "fmt"

func main() {
	fmt.Println(printSquare(square,3))

}
func printSquare(f func(int) int, a int) string {
	x := a * a
	return fmt.Sprintf("the numer %v is squared to be %v",a,x)
}

func square(n int) int {
	return n * n
}