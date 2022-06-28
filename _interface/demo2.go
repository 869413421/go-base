package main

type A interface {
	Foo()
}

type B interface {
	A
	Bar()
}

type T struct {
}

func (t T) Foo() {

}

func (t T) Bar() {

}

func TestA(a A) {

}

func TestB(b B) {

}

func main() {
	t := new(T)
	//t中实现了Foo方法所以可以存入
	TestA(t)
	//t中实现了Foo和Bar方法所以可以存入
	TestB(t)
}
