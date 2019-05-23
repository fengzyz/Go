package main

import (
	"fmt"
)

const (
	A int = 9
	B int = 10
)

func main()  {
	a := A + B
	fmt.Println(a)
	numbers := [5]int{1, 2, 3, 5}
	for i,x := range numbers{
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	x,y := 15,30

	fmt.Printf("交换前 x 的值为 : %d\n", x )
	fmt.Printf("交换前 y 的值为 : %d\n", y )
	swap(&x,&y)
	//allPrime()
	fmt.Printf("交换后 x 的值 : %d\n", x )
	fmt.Printf("交换后 y 的值 : %d\n", y )
}
//num以内所有素数
func allPrime(){
	var flag bool
	count := 1
	for count < 100 {
		count ++
		flag = true
		for tmp := 2;tmp < count;tmp++ {
			if count % tmp == 0{
				flag = false
			}
		}
		if flag{
			fmt.Println(count,"素数")
		}else {
			 continue
		}
	}
}
//交换两个值
func swap(x *int,y *int) int {
	var temp int
	temp = *x
	*x = *y
	*y = temp
	return temp
}
