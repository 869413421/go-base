package main

import (
	"fmt"
	"reflect"
)

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		case reflect.Bool:
			fmt.Println(arg, "is an bool type.")
		case reflect.Slice:
			fmt.Println(arg, "is an Slice type.")

		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	myPrintf(1, "1", [1]int{1}, true,[]string{"111"})
}
