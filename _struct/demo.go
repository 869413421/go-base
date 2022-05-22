package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float32
}

func NewStudent(id uint, name string, male bool, score float32) *Student {
	return &Student{
		id:    id,
		name:  name,
		male:  male,
		score: score,
	}
}

func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s Student) String() string {
	return fmt.Sprintf("{id1:%d,name:%s,male:%t,score:%f}", s.id, s.name, s.male, s.score)
}

func main() {
	student := NewStudent(1, "xiaoming", true, 100)
	student.SetName("test")
	fmt.Println(student.GetName())
	fmt.Println(student)
}
