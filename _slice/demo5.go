package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = slice1[:len(slice1)-5] // 删除尾部的5个元素
	fmt.Println(slice1)
	slice1 = slice1[1:] // 删除头部的1个元素
	fmt.Println(slice1)

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice4 := append(slice3[:0], slice3[3:]...) //删除开头三个元素
	//fmt.Println(slice4)

	//slice5 := append(slice3[:1], slice3[4:]...) // 删除中间三个元素
	//fmt.Println(slice5)

	//slice6 := append(slice3[:0], slice3[:7]...)  // 删除最后三个元素
	//fmt.Println(slice6)

	slice7 := slice3[:copy(slice3, slice3[3:])]  // 删除开头前三个元素
	fmt.Println(slice7)
}
