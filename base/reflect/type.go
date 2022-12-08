package reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

type Stu struct {
}

type Tea struct {
	Stu
}

func (*Stu) ToString() {}

func (*Tea) Test()    {}
func (Tea) SayHello() {}

type UStu struct {
	Username string
	age      int
}

type UAdmin struct {
	UStu
	title string
}

type UserModel struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
}

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

type Data struct {
	b byte
	x int32
}

func (*Data) String() string {
	return ""
}

type IData struct {
}

func (*IData) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*IData) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}

	return fmt.Sprintf(s, c)
}
