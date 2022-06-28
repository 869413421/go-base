package main

import (
	"fmt"
	"reflect"
)

type IAnimal interface {
	Call(s string)
}

type Dog struct {
}

func (d Dog) Call(s string) {
	fmt.Println(s)
}

func main() {
	dog := Dog{}
	var ianimal IAnimal = &dog
	if animal, ok := ianimal.(IAnimal); ok {
		animal.Call("ssss")
	} else {
		fmt.Println("xxxx")
	}

	fmt.Println(reflect.TypeOf(ianimal))
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("is int")
	}

	a := i.(int)
	fmt.Println(a)

	dog2 := Dog{}
	var f interface{} = dog2
	if dog3, ok := f.(Dog); ok {
		fmt.Println(dog3)
	}

}
