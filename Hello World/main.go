package main

import (
	"fmt"
	"net/http"
)
const partNumber=":8080"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is Home Page")
}
func About(w http.ResponseWriter, r *http.Request){
	sum:=addValues(2,2)
	fmt.Fprintf(w,fmt.Sprintf("About page with 2 + 2 equals to %d ",sum))
}
func addValues(x,y int) int{
	return x+y
}

func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	
	_=http.ListenAndServe(portNumber,nil)
}