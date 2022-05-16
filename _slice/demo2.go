package main

import "fmt"

func main() {
	// 创建一个数组
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组片段创建切片
	slice1 := nums[0:3]
	slice2 := nums[5:9]

	// 基于数组全部创建切片
	sliceAll := nums[:]

	// 基于切片创建切片
	slice3 := slice1[:2]

	// 直接创建,创建初始长度为10的切片。
	slice4 := make([]int, 10)

	// 直接创建,创建初始长度为10，容量为20的切片。
	slice5 := make([]int, 10, 20)

	// 还可以直接创建并初始化包含 5 个元素的数组切片
	slice6 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(sliceAll)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5, cap(slice5))
	fmt.Println(slice6, cap(slice6))

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i, v := range nums {
		fmt.Println(i, v)
	}
}
