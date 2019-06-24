package main

import (
	"fmt"
	"math"
	"strconv"
)

func main()  {

	var str string = "abc"
	var newS string
	fmt.Println("The size of ints is: %d\n",strconv.IntSize)
	an, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("str %s is not an integer - exiting with error\n", str)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	t, ok := mySqrt(25.0)
	if ok { fmt.Println(t) }
}

func mySqrt(f float64) (c float64,ok bool) {

	if f < 0 {
		return
	}
	return math.Sqrt(f),true
}
