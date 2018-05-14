package main

import "fmt"

//go 可变参数
func sumAgrs(valus ...int) int {
	sum := 0;
	for i := range valus{
		sum += valus[i]
	}
	return sum
}

type Book struct {
	name string
	title string
	id    int
	author string
	cotent string
}



func main() {
	fmt.Println(sumAgrs(1,8,9,10))
	var book Book
	book.title = "facebook"
	fmt.Println(book.title)
}
