package cgo

import "fmt"

type ZBuf struct {
	b *Buf
}

func (zb *ZBuf) Clear() {
	if zb.b != nil {
		NewPool().Revert(zb.b)
		zb.b = nil
	}
}

func (zb *ZBuf) Pop(len int) {
	if zb.b == nil || len > zb.b.Length() {
		return
	}

	zb.b.Pop(len)

	if zb.b.Length() == 0 {
		NewPool().Revert(zb.b)
		zb.b = nil
	}
}

func (zb *ZBuf) Data() []byte {
	if zb.b == nil {
		return nil
	}

	return zb.b.GetBytes()
}

func (zb *ZBuf) Adjust() {
	if zb.b == nil {
		zb.b.Adjust()
	}
}

func (zb *ZBuf) Read(src []byte) (err error) {
	if zb.b == nil {
		zb.b, err = NewPool().Alloc(len(src))
		if err != nil {
			fmt.Println("pool alloc error ", err)
			return err
		} else {
			if zb.b.Head() != 0 {
				return nil
			}

			if zb.b.Capacity-zb.b.Length() < len(src) {
				NewBuf, err := NewPool().Alloc(len(src) + zb.b.Length())
				if err != nil {
					fmt.Println(err)
					return err
				}

				NewBuf.Copy(zb.b)
				NewPool().Revert(zb.b)
				zb.b = NewBuf
			}
		}
	}

	zb.b.SetBytes(src)
	return nil
}
