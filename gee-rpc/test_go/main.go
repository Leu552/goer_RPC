package main

import "fmt"

func main() {
	data := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data {
		v.num *= 10 // 直接使用指针更新
	}
	fmt.Println(data[0], data[1], data[2]) // &{10} &{20} &{30}
	fmt.Printf("%p\n %p\n %p\n", data[0], data[1], data[2])
	//	输出data元素的类型和值
	fmt.Printf("%T %v\n", data, *data[0]) // []*struct { num int } {&{10}, &{20}, &{30}}

}
