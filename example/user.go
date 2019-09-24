package main

import "fmt"


func main(){
	data , i := [3]int{1,2,3}, 0
	i, data[i] = 500, 10
	fmt.Println(i, data)
}

