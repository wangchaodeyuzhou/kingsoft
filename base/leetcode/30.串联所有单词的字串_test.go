package leetcode

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo", "barr", "wing", "ding", "wing"}
	substring := findSubstring(s, words)
	fmt.Println(substring)
}

func findSubstring(s string, words []string) []int {
	res, strLen, wordLen := []int{}, len(s), len(words[0])
	for i := 0; i < wordLen; i++ {
		// 边界条件判断
		if i+wordLen*len(words) > strLen {
			break
		}
		// 初始化 differ, 分割字串
		differ := map[string]int{}
		for j := 0; j < len(words); j++ {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			differ[word]++
		}
		// 计算 滑动窗口 与 word 的 差值
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		// 开始滑动窗口
		for start := i; start < strLen-wordLen*len(words)+1; start += wordLen {
			if start != i {
				// 向右滑动窗口
				word := s[start+(len(words)-1)*wordLen : start+len(words)*wordLen]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}

				// 向左滑动窗口
				word = s[start-wordLen : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}

			if len(differ) == 0 {
				res = append(res, start)
			}
		}
	}

	return res
}
