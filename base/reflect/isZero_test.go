package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Personal struct {
	Name string
}

func TestIsZero(t *testing.T) {
	var p Personal
	p1 := &Personal{
		Name: "test",
	}
	fmt.Println(reflect.ValueOf(p).IsZero())
	fmt.Println(reflect.ValueOf(p1).IsZero())
	fmt.Println(reflect.DeepEqual(p, Personal{}))
	fmt.Println(reflect.DeepEqual(p1, &Personal{}))
}
