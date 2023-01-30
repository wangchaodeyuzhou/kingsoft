package interview

import (
	"fmt"
	"syscall"
	"testing"
	"unsafe"
)

func TestHelloWorld(t *testing.T) {
	// dll, err := syscall.LoadDLL("dllTest.dll")
	// if err != nil {
	// 	panic(err)
	// }
	// set go env GOARCH = 386

	dll := syscall.NewLazyDLL("dllTest.dll")
	hello := dll.NewProc("HelloWord")
	hello.Call()

	fmt.Println("========")
	sum := dll.NewProc("AddNums")
	call, _, _ := sum.Call(6, 2)
	fmt.Printf("%T\n", call)
	pointer := *((*int32)(unsafe.Pointer(&call)))
	fmt.Println(pointer)

}
