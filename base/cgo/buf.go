package cgo

import "C"
import (
	"fmt"
	"unsafe"
)

type Buf struct {
	Next     *Buf
	Capacity int
	length   int
	head     int
	data     unsafe.Pointer
}

func NewBuf(size int) *Buf {
	return &Buf{
		Capacity: size,
		length:   0,
		head:     0,
		Next:     nil,
		data:     Malloc(size),
	}
}

func (b *Buf) SetBytes(src []byte) {
	Memcpy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), src, len(src))
	b.length += len(src)
}

func (b *Buf) GetBytes() []byte {
	data := C.GoBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), C.int(b.length))
	return data
}

func (b *Buf) Copy(other *Buf) {
	Memcpy(b.data, other.GetBytes(), other.length)
	b.head = 0
	b.length = other.length
}

func (b *Buf) Pop(len int) {
	if b.data == nil {
		fmt.Printf("pop data is nil")
		return
	}
	if len > b.length {
		fmt.Printf("pop len > length")
		return
	}
	b.length -= len
	b.head += len
}

func (b *Buf) Adjust() {
	if b.head != 0 {
		if b.length != 0 {
			Memmove(b.data, unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), b.length)
		}
		b.head = 0
	}
}

func (b *Buf) Clear() {
	b.length = 0
	b.head = 0
}

func (b *Buf) Head() int {
	return b.head
}

func (b *Buf) Length() int {
	return b.length
}
