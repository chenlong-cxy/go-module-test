package model

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) Student {
	return Student{
		Name: name,
		Age:  age,
	}
}
