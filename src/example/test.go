package main

import "go_learn/src/simplemath"
import "fmt"

func main(){
	a,b := 20,17
	var name  = simplemath.Add(8,10)
	fmt.Println(name)
	fmt.Printf("a= %b\n",a)
	fmt.Printf("b= %b\n",b)
	c := a & b
	fmt.Printf("c= %b\n",c)

	str := "Hello,世界"
	fmt.Printf(str)
	n := len(str)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		ch := str[i]    // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}


	// 通过二维数组生成九九乘法表
	var multi [9][9]string

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 {  // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1 * n2)
		}
	}
	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2)  // 位宽为8，左对齐
		}
		fmt.Println()
	}
}


