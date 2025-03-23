package main

import "fmt"

func main() {
	x:=foo()
	fmt.Println(x)

	i,s:=bar()
	fmt.Println(i,s)
}

func foo() int{
	
	return 50
}
func bar()(int ,string){
	return 100,"Rohit"
}