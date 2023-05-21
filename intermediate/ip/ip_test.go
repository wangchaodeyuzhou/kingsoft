package ip

import (
	"fmt"
	"testing"
)

func TestIP(t *testing.T) {
	p, _ := GetIPInternal()
	fmt.Println(p)
}
