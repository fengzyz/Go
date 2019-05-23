package main

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	num := []int{2, 15, 6, 9, 20, 8, 17, 5, 30}
	sum := int(0)
	for _,v := range num{
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(convertToBin(8),convertToBin(13),convertToBin(0),fibonacci1(4,1,1),bubble(num))
	
	i,j := 50,2014
	
	p := &i
	fmt.Println(*p)
	*p = 21 
	fmt.Println(i)
	
	p = &j
	*p = 5 
	fmt.Println(*p)
	fmt.Println(j)
	
	v := Vertex{10, 6}
	v.X = 4
	f := &v
	f.X = 1e9
	fmt.Println(v)
	fmt.Println(v.X)
	
}
//冒泡排序
func bubble(num []int) []int{
	flag := true
	for i := 1 ;i < len(num); i++ {
		for j := 0; j < len(num) - i ; j++ {
			if num[j] > num[j + 1]{
				temp := num[j + 1]
				num[j + 1] = num[j]
				num[j] = temp
				flag = false
			}
		}
		if flag{
			break
		}
	}
	return num
}

//选择排序


//获取整数的二进制
func convertToBin(n int)  string {
	result := ""
	if n == 0 {
		return strconv.Itoa(0) + result
	}
	for ; n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
		
	}
	return result
}

//golang实现递归
func fibonacci(n int) int {
	if n == 1 || n == 2{
		return 1
	}else{
		return fibonacci(n - 2) + fibonacci(n-1)
	}
}

//递归优化
func fibonacci1(n int, a int, b int) int {
	if n == 1{
		return a
	}else if n == 2{
		return b
	}else{
		return fibonacci1(n - 1, b, a + b)
	}
}
