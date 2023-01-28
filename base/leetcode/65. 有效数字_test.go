package leetcode

import (
	"fmt"
	"testing"
)

// 状态机 根据状态图写代码就行 https://leetcode.cn/problems/valid-number/solutions/564188/you-xiao-shu-zi-by-leetcode-solution-298l/
func TestIsNumber(t *testing.T) {
	fmt.Println(isNumber("1e.66"))
}

type State int
type CharType int

const (
	State_INITIAL State = iota
	State_Sign
	State_Int_Point
	State_Integer
	State_Without_Int_Point
	State_Fraction
	State_Exp
	State_Exp_Sign
	State_Exp_Number
)

const (
	Char_Number CharType = iota
	Char_Point
	Char_Exp
	Char_Sign
	Char_Illegal
)

func toCharType(s byte) CharType {
	switch s {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return Char_Number
	case 'e', 'E':
		return Char_Exp
	case '.':
		return Char_Point
	case '+', '-':
		return Char_Sign
	default:
		return Char_Illegal
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		State_INITIAL: {
			Char_Number: State_Integer,
			Char_Point:  State_Without_Int_Point,
			Char_Sign:   State_Sign,
		},
		State_Sign: {
			Char_Number: State_Integer,
			Char_Point:  State_Without_Int_Point,
		},
		State_Without_Int_Point: {
			Char_Number: State_Fraction,
		},
		State_Integer: {
			Char_Number: State_Integer,
			Char_Exp:    State_Exp,
			Char_Point:  State_Int_Point,
		},
		State_Int_Point: {
			Char_Exp:    State_Exp,
			Char_Number: State_Fraction,
		},
		State_Fraction: {
			Char_Number: State_Fraction,
			Char_Exp:    State_Exp,
		},
		State_Exp: {
			Char_Number: State_Exp_Number,
			Char_Sign:   State_Exp_Sign,
		},
		State_Exp_Sign: {
			Char_Number: State_Exp_Number,
		},
		State_Exp_Number: {
			Char_Number: State_Exp_Number,
		},
	}

	state := State_INITIAL
	for i := 0; i < len(s); i++ {
		ch := toCharType(s[i])
		if _, ok := transfer[state][ch]; !ok {
			return false
		} else {
			state = transfer[state][ch]
		}
	}

	return state == State_INITIAL || state == State_Fraction || state == State_Integer || state == State_Int_Point || state == State_Exp_Number
}
