package main

import (
	"fmt"
	"html/template"
	"net/http"
	
)

// const port= ":8080"
// func Home(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w,"This is the home Page")
// }

// func About(w http.ResponseWriter, r *http.Request){
// 	sum:=Add(2,5)
// 	_, _=fmt.Fprintf(w,fmt.Sprintf("This gives the sum = %d",sum))
// }
// func Add(x,y int) int{
// 	return x+y
// }

//handle errors
// func Divide(w http.ResponseWriter,r *http.Request){
// 	f, err:=divideValues(100.0,10.0)
// 	if err!=nil{
// 		fmt.Fprintf(w,"Cannot divide by zero")
// 	}
// 	fmt.Fprintf(w,fmt.Sprintf("%f divided by %f is %f",100.0,10.0,f))
// }
// func divideValues(x,y float32) (float32,error){
// 	if y<=0{
// 		err:=errors.New("cannot be divided by zero")
// 		return 0,err
// 	}
// 	result:=x/y
// 	return result,nil
// }


const port= ":8080"
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"home.page.html")
}

func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"about.page.html")
}
func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _:=template.ParseFiles("./templates/"+ tmpl)
	err:=parsedTemplate.Execute(w,nil)
	if err!=nil{
		fmt.Println("error parsing template",err)
		return
	}
}



func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	// http.HandleFunc("/divide",Divide)
	fmt.Println(fmt.Sprintf("Starting App at port: %s",port))
	_=http.ListenAndServe(port,nil)

}