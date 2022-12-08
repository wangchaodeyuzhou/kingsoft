package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestTestPointer(t *testing.T) {
	u := stu{name: "csd", add: 100}
	// p := Per{name: "cds"}
	fmt.Println(u.name)

	p := unsafe.Pointer(&u)
	px := (*Per)(p)
	px.name = "dcsd"
	fmt.Printf("%v\n", px)
}

func TestTestPointer2(t *testing.T) {
	slice := []int{1, 3, 4, 5, 6}
	for v := range slice {
		fmt.Println(v)
	}

	x := 0
	switch x {
	case 0:
		fmt.Println("0")
		fallthrough
	case 10:
		fmt.Println("10")
	case 90:
		fmt.Println("90")
	default:
		fmt.Println("default")
	}
}

type UserH struct {
	id   int
	name string
}

func (u *UserH) TestPionter() {
	fmt.Printf("TestPointer: %p, %v\n", u, u)
}

func (u UserH) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &u, u)
}

func TestTestPointer3(t *testing.T) {
	u := UserH{1, "Tom"}
	fmt.Printf("UserH: %p, %v\n", &u, u)

	mv := UserH.TestValue
	mv(u)

	mp := (*UserH).TestPionter
	mp(&u)

	mp2 := (*UserH).TestValue
	mp2(&u)
}

func TestTestPointer4(t *testing.T) {
	u := UserH{1, "tom"}

	var vi, pi any = u, &u
	// vi.(UserH).name = "Jack"
	pi.(*UserH).name = "jack"

	fmt.Printf("%v\n", vi.(UserH))
	fmt.Printf("%v\n", pi.(*UserH))
}

type iface struct {
	itab, data uintptr
}

func TestTestPointer5(t *testing.T) {
	var a any = nil
	var b any = (*int)(nil)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
}

type Request struct {
	data []int
	ret  chan int
}

func NewRequest(data ...int) *Request {
	return &Request{data: data, ret: make(chan int, 1)}
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}

	req.ret <- x
}

func TestTestPointer6(t *testing.T) {
	req := NewRequest(10, 30, 40)
	Process(req)
	fmt.Println(<-req.ret)
}
