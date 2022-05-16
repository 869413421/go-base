package main

const Pi float64 = 3.14159265358979323846
const zero = 0.0 // 无类型浮点常量
const (          // 通过一个 const 关键字定义多个常量，和 var 类似
	size int64 = 1024
	eof        = -1 // 无类型整型常量
)
const u, v float32 = 0, 3   // u = 0.0, v = 3.0，常量的多重赋值
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量

const d, e, f = "1", 1, 1.00
