package main

import "fmt"

func main() {
	x:=[]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(sum(x...))
	

	
	fmt.Println(bar(x))
	


}

func sum(ii ...int)  int{
	t:=0
	for _,v:=range ii{
		t+=v
	}
	return t;
	
}

func bar(ii []int)  int{
	t:=0
	for _,v:=range ii{
		t+=v
	}
	return t;
	
}