package _map

import (
	"fmt"
	"testing"
)

type Stu struct {
	name string
}

func TestMapLoop(t *testing.T) {
	sessionToId := make(map[Stu][]string)
	sessionToId[Stu{name: "1"}] = []string{"1", "11"}
	sessionToId[Stu{name: "2"}] = []string{"2", "22"}
	sessionToId[Stu{name: "3"}] = []string{"3", "33"}

	ss := make([]*Stu, 0)
	for stu, session := range sessionToId {
		u := stu
		ss = append(ss, &u)
		fmt.Println(session)
	}

	for _, s := range ss {
		fmt.Println(s)
	}
}
