package main

import "fmt"

func main() {
	var a float32
	var b float32
	var c float32
	var d float32
	var e float32
	var g float32 = 6480
	var j float32 = 18800
	var m float32 = 5930
	a=35280
	b=2600
	c=7200
	d=b/c
	e = a*d
	f:=b*3
	h:=g*d
	i:=h+m
	k:=j-i
	n:=k/4

	fmt.Printf("a: %f ,b:%f, c:%f , d:%f ,e:%f ,f:%f,h:%f,i:%f,k:%f,n:%f",a,b,c,d,e,f,h,i,k,n)
}
