package options

import (
	"fmt"
	"testing"
	"time"
)

func TestWithOptionTimeout(t *testing.T) {
	obj, err := NewClient(WithURL("hello"), WithOptionTimeout(50*time.Second))
	if err != nil {
		return
	}

	fmt.Printf("%v\n", obj)
	fmt.Println(obj.url, obj.optionTimeout)

}
