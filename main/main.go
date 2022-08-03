package main

import (
	"fmt"
	"go-project/service"
	"go-project/structs/user"
)

var (
	name1 string
	age1  int
)

func main() {
	//方法调用,结构体使用
	var yangmi user.Girl
	yangmi.Name = "杨幂"
	yangmi.Age = 35
	yangmi.Sex = "female"
	guanxiaotong := user.Girl{"关晓彤", 4, "female"}
	whoName := service.CompareAge(yangmi, guanxiaotong)
	fmt.Println(whoName)

	//遍历数组
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	//指针
	money := [3]int{10, 20, 30}
	var ptr [len(money)]*int
	for i := 0; i < len(money); i++ {
		ptr[i] = &money[i]
	}
	for i := 0; i < len(ptr); i++ {
		fmt.Println(*ptr[i])
	}

	//切片
	service.Slice()

	//实现接口
	service.InterfaceTest()

	//range范围,map
	service.MapTest()
	var testMap = make(map[string]string)
	testMap["name"] = "成都"
	fmt.Println(testMap["name"])

}
