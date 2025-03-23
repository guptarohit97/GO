

package main

import "fmt"

type Person struct {
	first string
	age   int
}

func Speak(p Person) {

	fmt.Println("My name is ", p.first, "and age is ",p.age)

}
func main(){
	p1:=Person{
		first:"Rohit",
		age:28,
	}
	Speak(p1)
}

