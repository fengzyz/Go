package main

import (
	"fmt"
)

func main(){
	fmt.Println("main() started")

	c := make(chan string)

	a := make(chan  int)

	go greet(c)

	go squaras(a)
	for val := range a{
		fmt.Println(val)
	}
	c <- "John"
	close(c)
	c <- "Mike"
	fmt.Println("main() stoped")

}
func greet(c chan string){
	fmt.Println("hello " + <- c + "！")
}

func squaras(c chan int){
	for i := 0; i < 9; i++  {
		c <- i * i
	}
}
