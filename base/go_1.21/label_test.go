package go_1_21

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8"}

HEAD:
	for _, k := range a {
		if k == "7" {
			fmt.Println("continue")
			continue HEAD
		}
		fmt.Println(k)
	}

	fmt.Println("end...")
}
