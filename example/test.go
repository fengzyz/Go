package main

import "fmt"

var (
	a int = 4
	b int = 5
)

//var c func(string) int
var e,f = 12,"abc"
func main()  {
	getName("Jack")
}

func getName(name string)  {
	d := 8
	e = a + d
	g := &a
	a = 20
	//*g = 8
	numb,numa,strs := numbers()
	fmt.Println(e, f, *g, numb, strs, numa, name)
}
func numbers()(int,string,string){
	a , b , c := 1 , "abc" , "str"
	return a,b,c
}