package main

import (
	fm "fmt"
	"os"
	"strings"
)


const  (
	SUNDAY  = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
	GOPATH = os.Getenv("GOPATH")
)

var (
	age int
	big bool
	str string
	ip *int
)
func main() {
	//t := test(SUNDAY,THURSDAY)


	age = 25
	big = true
	str := "The quick brown fox jumps over the lazy dog"
	strings.HasPrefix(str, "so")
	if(strings.HasPrefix(str,"so")){
		fm.Println("1222")
	}
	s1 := strings.Fields(str)
	fm.Println(s1)
	for _, val := range s1 {
		fm.Printf("%s - ", val)
	}
}
func test(a,b int) int  {
	return  a + b
}



