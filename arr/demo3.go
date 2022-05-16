package main

import "fmt"

func main() {
	var milti [9][9]string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 {
				continue
			}
			milti[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1 * n2)
		}
	}

	// 打印九九乘法表
	for _, v1 := range milti {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2)  // 位宽为8，左对齐
		}
		fmt.Println()
	}
}
