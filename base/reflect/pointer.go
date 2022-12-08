package reflect

import (
	"fmt"
	"unsafe"
)

type stu struct {
	name string
	add  int
}

type Per struct {
	name string
}

func TestPointer() {
	u := stu{name: "csd", add: 100}
	// p := Per{name: "cds"}
	fmt.Println(u.name)

	p := uintptr(unsafe.Pointer(&u))
	p += unsafe.Offsetof(u.add)
	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 900
	fmt.Printf("%v\n", u)

	slice := []int{1, 3, 4, 5, 6}
	for v := range slice {
		fmt.Println(v)
	}
}
