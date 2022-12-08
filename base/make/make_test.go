package make

import (
	"fmt"
	"testing"
)

func TestIMake(t *testing.T) {
	var makeints func(int, int) []int
	var makestrings func(int, int) []string

	IMake(Int, &makeints)
	IMake(String, &makestrings)

	x := makeints(5, 10)
	fmt.Printf("%#v\n", x)

	s := makestrings(3, 10)
	fmt.Printf("%#v\n", s)

}
