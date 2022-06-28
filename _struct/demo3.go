package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) Call() string {
	return "动物的叫声"
}

func (a Animal) FavorFood() string {
	return "爱吃的食物"
}

func (a Animal) GetName() string {
	return a.name
}

type Dog struct {
	*Animal
}

func (d Dog) FavorFood() string {
	return "骨头"
}

func (d Dog) Call() string {
	return "汪汪汪"
}

type Puppy struct {
	Animal
	Dog
}

func main() {
	animal := Animal{"狗"}
	dog := Dog{&animal}
	fmt.Println(dog.GetName(), "叫声:", dog.Call(), "喜爱的食物:", dog.FavorFood())
	fmt.Println(dog.Animal.name, "叫声:", dog.Animal.Call(), "喜爱的食物:", dog.Animal.FavorFood())
}
