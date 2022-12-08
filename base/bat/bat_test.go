package bat

import (
	"fmt"
	"testing"
)

func TestBattleType_GetHelloType(t *testing.T) {
	fmt.Println("cdsdcds ")
}
func NewBattleType() *BattleType {
	return &BattleType{Hello: &HelloType{say: nil}}
}

func TestBattleHH(t *testing.T) {
	//b := NewBattleType()
	//fmt.Printf("%v\n", b)

	b := &BattleType{Hello: nil}
	fmt.Printf("%v\n", b)

	g := b.GetHelloType()
	fmt.Printf("%v\n", g)
	gg := g.GetSayType()
	fmt.Printf("%v\n", gg)
	ggg := gg.GetTeam()
	fmt.Printf("%v\n", ggg)

	//fmt.Println("bat", "")
}
