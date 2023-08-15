package genericity

import (
	"fmt"
	"testing"
)

func TestGeneric(t *testing.T) {
	var m = map[int]string{1: "2", 2: "1", 6: "3"}
	fmt.Println(MapKeys[int, string](m))

	fmt.Println(MapKeys(m))
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(2)
	lst.Push(67)
	lst.Push(432)
	fmt.Println(lst.GetAll())
}
