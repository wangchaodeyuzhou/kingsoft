package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	fmt.Println(fullJustify(words, maxWidth))
}

func fullJustify(words []string, maxWidth int) []string {
	right, n := 0, len(words)
	ans := []string{}
	for {
		left := right
		sumLen := 0 // count 这 line 的 sum total
		// 统计一行最多放多少单词
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		// 结束条件
		if right == n {
			s := strings.Join(words[left:], " ")
			ans = append(ans, s+blank(maxWidth-len(s)))
			return ans
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		// 如果一行只有一个单词,
		if numWords == 1 {
			// 直接在后面进行填充空格
			ans = append(ans, words[left]+blank(numSpaces))
			continue
		}

		// 不止一个单词在一行中
		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)
		s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1))
		s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))
		ans = append(ans, s1+blank(avgSpaces)+s2)
	}
}

func blank(n int) string {
	return strings.Repeat(" ", n)
}
