package main

import (
	"encoding/json"
	"fmt"
)

type cb func(int) int

type User struct {
	Id int
	Sex int
	Name string
	Password string
}

func main(){
	fmt.Println();
	testcallBack(8,callBack)
	testcallBack(9, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return  x
	})
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	add_func := add(1,2)
	fmt.Println(add_func(8,5))

	var user  User
	//user.sex = 1
	user.getUserInfo()

}
func (c User) getAge(sex int) int{
	var age = 0
	if sex == 1 {
		age = 100
	}else if sex == 2{
		age = 85
	}
	return age
}
func (c User) getUserInfo() {
	c.Id = 8
	c.Name = "wangwu"
	c.Password = "123456"
	c.Sex = 1
	fmt.Println( c )
	data, _ := json.Marshal(c)
	fmt.Println(string(data))
}

func testcallBack(x int,f cb)  {
	f(x)
}
func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

func getSequence() func()int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
// 闭包使用方法
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
	i := 5
	return func(x3 int,x4 int) (int,int,int){
		i++
		return i,x1+x2,x3+x4
	}
}

