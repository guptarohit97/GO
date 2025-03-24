package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "Rohit",
	}
	fmt.Println(p)
	p=changeName(p,"Meena")
	fmt.Println(p)


	changeNamePtr(&p,"Meena")
	fmt.Println(p)
}
func changeName(p person,s string)person{
	p.first=s
	return p
}
func changeNamePtr(p *person,s string){
	p.first=s
}