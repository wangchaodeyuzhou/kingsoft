package http

import (
	"fmt"
	"net/textproto"
	"testing"
)

// 测试规范化 HTTP 头的键名: 通常使用规范化的键名。
// 规范化的键名是将 HTTP 头键名转换为特定的格式：每个单词的首字母大写，其余字母小写，并使用连字符 "-" 分隔不同的单词。
func TestMk(t *testing.T) {
	headerKey := textproto.CanonicalMIMEHeaderKey("username")
	fmt.Println(headerKey)
}
