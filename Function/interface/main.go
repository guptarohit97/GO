package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width float64
}
type circle struct{
	radius float64
}
func (s square)area()float64{
	return s.length*s.width
}
func (c circle)area()float64{
	return math.Pi*c.radius*c.radius
}
type shape interface{
	area() float64
}
func info(s shape) float64{
	return s.area()
}

func main(){
	c1:=circle{
		radius: 4,
	}
	s1:=square{
		length: 4,
		width: 4,
	}
	fmt.Println(info(c1))
	fmt.Println(info(s1))
}
