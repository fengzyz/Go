package main

import (
	"fmt"
	"math/cmplx"
	"math"
)
var (
	aa = 123
	bb = "ccc"
	dd = true
)
func variableZeroValue(){
	var a int       //变量定义 整形默认为0
	var b string    //字符串默认为空
	fmt.Printf("%d %q\n", a, b)
}

func variableInitalValue(){
	var a, b int = 3, 4
	var c string ="hello"
	fmt.Println(a,b,c)
}

func variableTypeValue(){
	var a, b, c = 3, 4, "111"
	fmt.Println(a,b,c)
}

/*
* :=只能用在方法内,第一次使用变量
* 使用:=后不能在使用var
*
 */
func variableSholter(){
	a, b, c, d := 3, 4, "111",true
	b = 5
	fmt.Println(a,b,c,d)
}
//求三角形斜边
func triangle(){
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b))) //go不支持隐式，所以必须强制转换
	fmt.Println(c)
}
//枚举用法，iota 实现自增
func enums(){
	const (
		cpp = iota
		java
		golang
		python
	)
	const (
		b = 1 << (10 * iota)
		kb
		gb
		db
		mb
	)
	fmt.Println(cpp,java,golang,python)
	fmt.Println(b, kb, gb, db, mb)
}
func euler(){
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi )+ 1)
	}

func main() {
	fmt.Println("Hello Wrold")
	variableZeroValue()
	variableInitalValue()
	variableTypeValue()
	variableSholter()
	fmt.Println(aa,bb,dd)
	euler()
	triangle()
	enums()
}
