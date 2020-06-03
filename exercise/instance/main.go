package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

//method형태
func (t *Student) SetName(newName string) {
	t.name = newName
}

//function 형태
func SetName(t *Student, newName string) {
	t.name = newName
}

func PrintStudent(u *Student) {
	fmt.Println(u)
}

func (t *Student) SetAge(age int) {
	t.age = age
}

func main() {
	var b *Student
	b = &Student{"BBB", 40, 50}
	a := Student{"aaa", 17, 90}

	// SetName(&a, "bbb")
	a.SetName("AAA") //oop 방식 메서드를 가지고있음
	a.SetAge(20)

	fmt.Println(a)
	fmt.Println(a.name)
	PrintStudent(&a)
	PrintStudent(b)
}
