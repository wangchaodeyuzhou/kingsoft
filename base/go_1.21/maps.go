package go_1_21

import (
	"fmt"
	"maps"
)

func mapsTest() {
	var mp = map[int]string{1: "23", 2: "33"}
	m := make(map[int]string, len(mp))
	maps.Copy(m, mp)
	fmt.Println(mp, m)
	clear(m)
	fmt.Println(mp, m)
}
