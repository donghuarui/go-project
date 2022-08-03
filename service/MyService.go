package service

import (
	"fmt"
	"go-project/structs/user"
)

func CompareAge(girl1 user.Girl, girl2 user.Girl) string {
	if girl1.Age > girl2.Age {
		return girl1.Name
	} else if girl1.Age < girl2.Age {
		return girl2.Name
	} else {
		return "一样大"
	}
}

func def(name1 string, name2 string) string {
	if name1 == "阿木木" {
		return name1
	} else if name2 == "阿木木" {
		return name2
	} else {
		return "没有阿木木"
	}
}

func Slice() {
	fmt.Println("====================切片================================")
	//切片
	var numbers []int
	numbers1 := append(numbers, 1, 2, 3)
	var numbers2 []int = make([]int, 2, 2)
	fmt.Println(numbers1)
	fmt.Println("长度=%d,最大容量=%d +", len(numbers1), cap(numbers1))
	copy(numbers2, numbers1)
	fmt.Println("numbers2=", numbers2)
}

func InterfaceTest() {
	//实现接口
	var fruiter FruitEr
	apple := Apple{}
	fruiter = apple
	fmt.Println(fruiter.Eat("好吃"))
	banana := new(Banana)
	fruiter = banana
	fmt.Println(fruiter.Eat("不好吃"))
}

func MapTest() {
	//range范围,map
	girl := map[string]string{"name": "赵丽颖", "age": "22", "sex": "female"}
	for key, value := range girl {
		fmt.Println("key:", key, "value:", value)
	}
}
