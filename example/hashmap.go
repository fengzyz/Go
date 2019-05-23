package main

import "fmt"

func main() {
	fmt.Printf("123")
}

//定义结构
type HashMap struct {
	key string
	value string
	hashCode  int
	next *HashMap
}
var table [16](*HashMap)

// 初始化table
func initTable()  {
	for i := range table{
		table[i] = &HashMap{"","",i,nil}
	}
}

func getInstance() [16]( *HashMap) {
	if(table[0] == nil){
		initTable()
	}
	return table
}

func genHashCode(k string) int  {
	if len(k) == 0 {
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) -1
	for i := range k{
		if i == lastIndex{
			hashCode += int(k[i])
			break
		}
		hashCode += (hashCode + int(k[i])*30)
	}
	return hashCode
}

func indexTable(hashCode int)int   {
	return hashCode%16
}

