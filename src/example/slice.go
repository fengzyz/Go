package main

import "fmt"

func main(){
	months := [...] string {
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}
	// 基于数组创建数组切片
	q2 := months[3:6]    // 第二季度
	summer := months[5:8]  // 夏季
	all := months[:]      //全年
	firsthalf := months[:6]
	fmt.Println(q2)
	fmt.Println(summer)
	fmt.Println(all)
	fmt.Println(firsthalf)
}
