package main

type Student struct {
	name string
	age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student)

	s.age = age
	s.name = name

	return s
}

func main() {
	StudentRegister("ji", 13)
}
